package streamdeck

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif" // Allow gifs to be loaded
	"image/jpeg"
	_ "image/png" // Allow pngs to be loaded
	"os"

	"github.com/disintegration/gift"
	"golang.org/x/image/bmp"
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

func getImageForButton(img image.Image, btnFormat string) ([]byte, error) {
	var b bytes.Buffer
	switch btnFormat {
	case "JPEG":
		jpeg.Encode(&b, img, nil)
	case "BMP":
		bmp.Encode(&b, img)
	default:
		return nil, errors.New("Unknown button image format: " + btnFormat)
	}
	return b.Bytes(), nil
}

func getSolidColourImage(colour color.Color) *image.RGBA {
	ButtonSize := 80
	img := image.NewRGBA(image.Rect(0, 0, ButtonSize, ButtonSize))
	//colour := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(colour), image.Point{0, 0}, draw.Src)
	return img
}

func getImageFile(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}
