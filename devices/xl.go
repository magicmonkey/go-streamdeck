package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

var (
	name                     string
	buttonWidth              uint
	buttonHeight             uint
	imageReportPayloadLength uint
)

// GetImageHeaderXl returns the USB comms header for a button image for the XL
func GetImageHeaderXl(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	thisLength := uint(0)
	if imageReportPayloadLength < bytesRemaining {
		thisLength = imageReportPayloadLength
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
	name = "Streamdeck XL"
	buttonWidth = 96
	buttonHeight = 96
	imageReportPayloadLength = 1024
	streamdeck.RegisterDevicetype(
		name, // Name
		image.Point{X: int(buttonWidth), Y: int(buttonHeight)}, // Width/height of a button
		0x6c,                     // USB productID
		[]byte{'\x03', '\x02'},   // Reset packet
		32,                       // Number of buttons
		[]byte{'\x03', '\x08'},   // Set brightness packet preamble
		4,                        // Button read offset
		"JPEG",                   // Image format
		imageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderXl,         // Function to get the comms image header
	)
}
