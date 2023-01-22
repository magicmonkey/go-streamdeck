package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

var (
	omk2Name                     string
	omk2ButtonWidth              uint
	omk2ButtonHeight             uint
	omk2ImageReportPayloadLength uint
)

// GetImageHeaderOv2 returns the USB comms header for a button image for the XL
func GetImageHeaderOMK2(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
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
	omk2Name = "Stream Deck MK.2"
	omk2ButtonWidth = 72
	omk2ButtonHeight = 72
	omk2ImageReportPayloadLength = 1024
	streamdeck.RegisterDevicetype(
		omk2Name, // Name
		image.Point{X: int(omk2ButtonWidth), Y: int(omk2ButtonHeight)}, // Width/height of a button
		0x80,                         // USB productID
		resetPacket32(),              // Reset packet
		15,                           // Number of buttons
		3,                            // Number of rows
		5,                            // Number of columns
		brightnessPacket32(),         // Set brightness packet preamble
		4,                            // Button read offset
		"JPEG",                       // Image format
		omk2ImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderOMK2,           // Function to get the comms image header
	)
}
