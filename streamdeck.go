package streamdeck

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"github.com/disintegration/gift"
	"github.com/karalabe/hid"
)

const VendorID = 4057
const ProductID = 0x6c

var Black = color.RGBA{0, 0, 0, 0}

type Device struct {
	fd *hid.Device
}

// Opens a new StreamdeckXL device, and returns a handle
func Open() *Device {
	devices := hid.Enumerate(VendorID, ProductID)
	if len(devices) == 0 {
		panic("no stream deck device found")
	}
	id := 0
	dev, err := devices[id].Open()
	if err != nil {
		panic(err)
	}
	retval := &Device{}
	retval.fd = dev
	return retval
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
		d.WriteColorToButton(Black, i)
	}
}

// Writes a specified color to the given button
func (d *Device) WriteColorToButton(colour color.Color, btnIndex int) {
	img := getSolidColourImage(colour)
	d.rawWriteToButton(btnIndex, getImageAsJpeg(img))
}

func getSolidColourImage(colour color.Color) *image.RGBA {
	ButtonSize := 96
	img := image.NewRGBA(image.Rect(0, 0, ButtonSize, ButtonSize))
	//colour := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(colour), image.Point{0, 0}, draw.Src)
	return img
}

// Writes a specified JPEG file to the given button
func (d *Device) WriteImageToButton(filename string, btnIndex int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}
	d.writeToButton(btnIndex, img)
	return nil
}

// Registers a callback to be called whenever a button is pressed
func (d *Device) ButtonPress(f func(int, *Device)) {
	var buttonMask [32]bool
	for {
		data := make([]byte, 50)
		_, err := d.fd.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < 32; i++ {
			if data[4+i] == 1 {
				if !buttonMask[i] {
					f(i, d)
				}
				buttonMask[i] = true
			} else {
				buttonMask[i] = false
			}
		}
	}
}

func (d *Device) writeToButton(btnIndex int, raw_img image.Image) error {
	img := ResizeAndRotate(raw_img, 96, 96)
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

// TODO make private
func ResizeAndRotate(img image.Image, width, height int) image.Image {
	g := gift.New(
		gift.Resize(width, height, gift.LanczosResampling),
		//gift.UnsharpMask(1, 1, 0),
		gift.Rotate180(),
	)
	res := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(res, img)
	return res
}

func getImageAsJpeg(img image.Image) []byte {
	var b bytes.Buffer
	jpeg.Encode(&b, img, nil)
	return b.Bytes()
}
