package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

func init() {
	streamdeck.RegisterDevicetype(
		"Streamdeck XL",           // Name
		image.Point{X: 96, Y: 96}, // Width/height of a button
		0x6c,                      // USB productID
		[]byte{'\x03', '\x02'},    // Reset packet
		32,                        // Number of buttons
		[]byte{'\x03', '\x08'},    // Set brightness packet preamble
		4,                         // Button read offset
	)
}
