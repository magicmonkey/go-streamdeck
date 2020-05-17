package streamdeck

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"

	"github.com/disintegration/gift"
)

func resizeAndRotate(img image.Image, width, height int) image.Image {
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

func getSolidColourImage(colour color.Color) *image.RGBA {
	ButtonSize := 96
	img := image.NewRGBA(image.Rect(0, 0, ButtonSize, ButtonSize))
	//colour := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(colour), image.Point{0, 0}, draw.Src)
	return img
}
