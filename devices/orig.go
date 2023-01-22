package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

var (
	originalName                     string
	originalButtonWidth              uint
	originalButtonHeight             uint
	originalImageReportPayloadLength uint
)

// GetImageHeaderMini returns the USB comms header for a button image for the original
func GetImageHeaderOriginal(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	var thisLength uint
	if originalImageReportPayloadLength < bytesRemaining {
		thisLength = originalImageReportPayloadLength
	} else {
		thisLength = bytesRemaining
	}
	header := []byte{
		'\x02',
		'\x01',
		byte(pageNumber + 1),
		0,
		get_header_element(thisLength, bytesRemaining),
		byte(btnIndex + 1),
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
		'\x00',
	}

	return header
}

func init() {
	originalName = "Stream Deck Original"
	originalButtonWidth = 72
	originalButtonHeight = 72
	originalImageReportPayloadLength = 8191 //8191
	streamdeck.RegisterDevicetype(
		originalName, // Name
		image.Point{X: int(originalButtonWidth), Y: int(originalButtonHeight)}, // Width/height of a button
		0x60,                             // USB productID
		resetPacket17(),                  // Reset packet
		15,                               // Number of buttons
		3,                                // Number of rows
		5,                                // Number of cols
		brightnessPacket17(),             // Brightness packet
		1,                                // Button read offset
		"BMP",                            // Image format
		originalImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderOriginal,           // Function to get the comms image header
	)
}
