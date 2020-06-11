package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

var (
	ov2Name                     string
	ov2ButtonWidth              uint
	ov2ButtonHeight             uint
	ov2ImageReportPayloadLength uint
)

// GetImageHeaderOv2 returns the USB comms header for a button image for the XL
func GetImageHeaderOv2(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	thisLength := uint(0)
	if ov2ImageReportPayloadLength < bytesRemaining {
		thisLength = ov2ImageReportPayloadLength
	} else {
		thisLength = bytesRemaining
	}
	header := []byte{'\x02', '\x07', byte(btnIndex)}
	if thisLength == bytesRemaining {
		header = append(header, '\x01')
	} else {
		header = append(header, '\x00')
	}

	header = append(header, byte(thisLength&0xff))
	header = append(header, byte(thisLength>>8))

	header = append(header, byte(pageNumber&0xff))
	header = append(header, byte(pageNumber>>8))

	return header
}

func init() {
	ov2Name = "Streamdeck (original v2)"
	ov2ButtonWidth = 72
	ov2ButtonHeight = 72
	ov2ImageReportPayloadLength = 1024
	streamdeck.RegisterDevicetype(
		ov2Name, // Name
		image.Point{X: int(ov2ButtonWidth), Y: int(ov2ButtonHeight)}, // Width/height of a button
		0x6d,                        // USB productID
		[]byte{'\x03', '\x02'},      // Reset packet
		15,                          // Number of buttons
		[]byte{'\x03', '\x08'},      // Set brightness packet preamble
		4,                           // Button read offset
		"JPEG",                      // Image format
		ov2ImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderOv2,           // Function to get the comms image header
	)
}
