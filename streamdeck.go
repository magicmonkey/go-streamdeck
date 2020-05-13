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

type Device struct {
	fd *hid.Device
}

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

func (d *Device) Close() {
	d.fd.Close()
}

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

func (d *Device) ClearButtons() {
	for i := 0; i < 32; i++ {
		d.WriteColorToButton(0, 0, 0, i)
	}
}

func (d *Device) WriteColorToButton(red uint8, green uint8, blue uint8, btnIndex int) {
	ButtonSize := 96
	img := image.NewRGBA(image.Rect(0, 0, ButtonSize, ButtonSize))
	color := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(color), image.Point{0, 0}, draw.Src)
	raw_image := getImageAsJpeg(img)
	d.writeToButton(btnIndex, raw_image)
}

func (d *Device) WriteImageToButton(filename string, btnIndex int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}
	img = resizeAndRotate(img, 96, 96)
	d.writeToButton(btnIndex, getImageAsJpeg(img))
	return nil
}

// Based on set_key_image from https://github.com/abcminiuser/python-elgato-streamdeck/blob/master/src/StreamDeck/Devices/StreamDeckXL.py#L151
func (d *Device) writeToButton(btnIndex int, raw_image []byte) error {
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

func (d *Device) ButtonPress(f func(int, *Device)) {
	for {
		data := make([]byte, 50)
		_, err := d.fd.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < 32; i++ {
			if data[4+i] == 1 {
				f(i, d)
			}
		}
	}
}

// resize returns a resized copy of the supplied image with the given width and height.
func resizeAndRotate(img image.Image, width, height int) image.Image {
	g := gift.New(
		gift.Resize(width, height, gift.LanczosResampling),
		gift.UnsharpMask(1, 1, 0),
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
