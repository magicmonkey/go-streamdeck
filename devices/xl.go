package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

var (
	xlName                     string
	xlButtonWidth              uint
	xlButtonHeight             uint
	xlImageReportPayloadLength uint
)

// GetImageHeaderXl returns the USB comms header for a button image for the XL
func GetImageHeaderXl(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	thisLength := uint(0)
	if xlImageReportPayloadLength < bytesRemaining {
		thisLength = xlImageReportPayloadLength
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
	xlName = "Streamdeck XL"
	xlButtonWidth = 96
	xlButtonHeight = 96
	xlImageReportPayloadLength = 1024
	streamdeck.RegisterDevicetype(
		xlName, // Name
		image.Point{X: int(xlButtonWidth), Y: int(xlButtonHeight)}, // Width/height of a button
		0x6c,                       // USB productID
		resetPacket32()[0:1],       // Reset packet
		32,                         // Number of buttons
		4,                          // Number of rows
		8,                          // Number of cols
		brightnessPacket32()[0:1],  // Set brightness packet preamble
		4,                          // Button read offset
		"JPEG",                     // Image format
		xlImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderXl,           // Function to get the comms image header
	)
}
