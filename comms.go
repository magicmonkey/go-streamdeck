package streamdeck

import (
	"errors"
	"image"
	"image/color"

	"github.com/karalabe/hid"
)

const VendorID = 4057
const ProductID = 0x6c

type Device struct {
	fd                   *hid.Device
	buttonPressListeners []func(int, *Device, error)
}

func Open() (*Device, error) {
	return rawOpen(true)
}

func OpenWithoutReset() (*Device, error) {
	return rawOpen(false)
}

// Opens a new StreamdeckXL device, and returns a handle
func rawOpen(reset bool) (*Device, error) {
	devices := hid.Enumerate(VendorID, ProductID)
	if len(devices) == 0 {
		return nil, errors.New("no stream deck device found")
	}
	id := 0
	dev, err := devices[id].Open()
	if err != nil {
		return nil, err
	}
	retval := &Device{}
	retval.fd = dev
	if reset {
		retval.ResetComms()
	}
	go retval.buttonPressListener()
	return retval, nil
}

// Closes the device
func (d *Device) Close() {
	d.fd.Close()
}

// Sets the button brightness
// pct is an integer between 0-100
func (d *Device) SetBrightness(pct int) {
	if pct < 0 {
		pct = 0
	}
	if pct > 100 {
		pct = 100
	}

	payload := []byte{'\x03', '\x08', byte(pct)}
	d.fd.SendFeatureReport(payload)
}

// Writes a black square to all buttons
func (d *Device) ClearButtons() {
	for i := 0; i < 32; i++ {
		d.WriteColorToButton(i, color.Black)
	}
}

// Writes a specified color to the given button
func (d *Device) WriteColorToButton(btnIndex int, colour color.Color) {
	img := getSolidColourImage(colour)
	d.rawWriteToButton(btnIndex, getImageAsJpeg(img))
}

// Writes a specified image file to the given button
func (d *Device) WriteImageToButton(btnIndex int, filename string) error {
	img, err := getImageFile(filename)
	if err != nil {
		return err
	}
	d.WriteRawImageToButton(btnIndex, img)
	return nil
}

func (d *Device) buttonPressListener() {
	var buttonMask [32]bool
	for {
		data := make([]byte, 50)
		_, err := d.fd.Read(data)
		if err != nil {
			d.sendButtonPressEvent(-1, err)
			break
		}
		for i := 0; i < 32; i++ {
			if data[4+i] == 1 {
				if !buttonMask[i] {
					d.sendButtonPressEvent(i, nil)
				}
				buttonMask[i] = true
			} else {
				buttonMask[i] = false
			}
		}
	}
}

func (d *Device) sendButtonPressEvent(btnIndex int, err error) {
	for _, f := range d.buttonPressListeners {
		f(btnIndex, d, err)
	}
}

// Registers a callback to be called whenever a button is pressed
func (d *Device) ButtonPress(f func(int, *Device, error)) {
	d.buttonPressListeners = append(d.buttonPressListeners, f)
}

func (d *Device) ResetComms() {
	payload := []byte{'\x03', '\x02'}
	d.fd.SendFeatureReport(payload)
}

func (d *Device) WriteRawImageToButton(btnIndex int, raw_img image.Image) error {
	img := resizeAndRotate(raw_img, 96, 96)
	return d.rawWriteToButton(btnIndex, getImageAsJpeg(img))
}

func (d *Device) rawWriteToButton(btnIndex int, raw_image []byte) error {
	// Based on set_key_image from https://github.com/abcminiuser/python-elgato-streamdeck/blob/master/src/StreamDeck/Devices/StreamDeckXL.py#L151
	page_number := 0
	bytes_remaining := len(raw_image)

	IMAGE_REPORT_LENGTH := 1024
	IMAGE_REPORT_HEADER_LENGTH := 8
	IMAGE_REPORT_PAYLOAD_LENGTH := IMAGE_REPORT_LENGTH - IMAGE_REPORT_HEADER_LENGTH

	payloads := make([][]byte, 3)

	for bytes_remaining > 0 {
		this_length := 0
		if IMAGE_REPORT_PAYLOAD_LENGTH < bytes_remaining {
			this_length = IMAGE_REPORT_PAYLOAD_LENGTH
		} else {
			this_length = bytes_remaining
		}
		bytes_sent := page_number * IMAGE_REPORT_PAYLOAD_LENGTH
		header := []byte{'\x02', '\x07', byte(btnIndex)}
		if this_length == bytes_remaining {
			header = append(header, '\x01')
		} else {
			header = append(header, '\x00')
		}

		header = append(header, byte(this_length&0xff))
		header = append(header, byte(this_length>>8))

		header = append(header, byte(page_number&0xff))
		header = append(header, byte(page_number>>8))

		payload := append(header, raw_image[bytes_sent:(bytes_sent+this_length)]...)
		padding := make([]byte, IMAGE_REPORT_LENGTH-len(payload))

		thingToSend := append(payload, padding...)
		d.fd.Write(thingToSend)
		payloads[page_number] = thingToSend

		bytes_remaining = bytes_remaining - this_length
		page_number = page_number + 1
	}
	return nil
}
