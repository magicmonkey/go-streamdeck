package streamdeck

import (
	"bytes"
	"errors"
	"fmt"
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

func resizeAndRotate(img image.Image, width, height int, devname string) image.Image {
	g, _ := deviceSpecifics(devname, width, height)
	res := image.NewRGBA(g.Bounds(img.Bounds()))
	g.Draw(res, img)
	return res
}

func deviceSpecifics(devName string, width, height int) (*gift.GIFT, error) {
	switch devName {
		case "Streamdeck XL", "Streamdeck (original v2)":
			return gift.New(
				gift.Resize(width, height, gift.LanczosResampling),
				gift.Rotate180(),
			), nil
		case "Streamdeck Mini":
			return gift.New(
				gift.Resize(width, height, gift.LanczosResampling),
				gift.Rotate90(),
				gift.FlipVertical(),
			), nil
		default:
			return nil, errors.New(fmt.Sprintf("Unsupported Device: %s", devName))
	}
}

func getImageForButton(img image.Image, btnFormat string) ([]byte, error) {
	var b bytes.Buffer
	switch btnFormat {
	case "JPEG":
		jpeg.Encode(&b, img, &jpeg.Options{Quality: 100})
	case "BMP":
		bmp.Encode(&b, img)
	default:
		return nil, errors.New("Unknown button image format: " + btnFormat)
	}
	return b.Bytes(), nil
}

func getSolidColourImage(colour color.Color, btnSize int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, btnSize, btnSize))
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
