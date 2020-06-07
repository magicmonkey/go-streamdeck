package devices

import (
	"image"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

func init() {
	streamdeck.RegisterDevicetype(
		"Streamdeck XL",
		image.Point{X: 96, Y: 96},
		0x6c,                   // productID
		[]byte{'\x03', '\x02'}, // reset packet
		32,
	)
}
