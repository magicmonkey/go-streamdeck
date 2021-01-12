package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

var (
	miniName                     string
	miniButtonWidth              uint
	miniButtonHeight             uint
	miniImageReportPayloadLength uint
	miniImageReportHeaderLength  uint
	miniImageReportLength        uint
)

// GetImageHeaderMini returns the USB comms header for a button image for the XL
func GetImageHeaderMini(bytesRemaining uint, btnIndex uint, pageNumber uint) []byte {
	thisLength := uint(0)
	if miniImageReportPayloadLength < bytesRemaining {
		thisLength = miniImageReportPayloadLength
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
	miniName = "Streamdeck Mini"
	miniButtonWidth = 80
	miniButtonHeight = 80
	miniImageReportLength = 1024
	miniImageReportHeaderLength = 16
	miniImageReportPayloadLength = miniImageReportLength - miniImageReportHeaderLength
	streamdeck.RegisterDevicetype(
		miniName, // Name
		image.Point{X: int(miniButtonWidth), Y: int(miniButtonHeight)}, // Width/height of a button
		0x63,                        // USB productID
		[]byte{'\x03', '\x00', '\x42', '\x00', '\x4C', '\x00', '\x33', '\x00', '\x31', '\x00', '\x4A', '\x00', '\x31', '\x00', '\x42', '\x00', '\x30', '\x00', '\x31', '\x00', '\x35', '\x00', '\x38', '\x00', '\x32', '\x00'},      // Reset packet
		6,                          // Number of buttons
		//[]byte{'\x05', '\x55', '\xaa', '\xd1', '\x01'},      // Set brightness packet preamble
		[]byte{'\x03', '\x00', '\x42', '\x00', '\x4C', '\x00', '\x33', '\x00', '\x31', '\x00', '\x4A', '\x00', '\x31', '\x00', '\x42', '\x00', '\x30', '\x00', '\x31', '\x00', '\x35', '\x00', '\x38', '\x00', '\x32', '\x00'},      // Reset packet
		0,                           // Button read offset
		"BMP",                      // Image format
		miniImageReportPayloadLength, // Amount of image payload allowed per USB packet
		GetImageHeaderMini,           // Function to get the comms image header
	)
}
