package streamdeck

import (
	"image"
	"image/color"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/font/gofont/gomedium"
)

type ButtonOptions struct {
	MarginX     int
	MarginY     int
	MaxTextSize float64
	AlignX      int
	AlignY      int
}

// WriteTextToButton is a low-level way to write text directly onto a button on the StreamDeck
func (d *Device) WriteTextToButton(btnIndex int, text string, textColor color.Color, backgroundColor color.Color) {
	img := GetImageWithText(text, textColor, backgroundColor, d.deviceType.imageSize.X)
	d.WriteRawImageToButton(btnIndex, img)
}

// WriteTextToButtonExt provides additional features like multiline text, margins and alignment
func (d *Device) WriteTextToButtonExt(btnIndex int, text string, textColor color.Color, backgroundColor color.Color, opt ButtonOptions) {
	img := GetImageWithTextExt(text, textColor, backgroundColor, d.deviceType.imageSize.X, opt)
	d.WriteRawImageToButton(btnIndex, img)
}

func testSizeWillFit(lines []string, linePadding int, btnSizeX int, btnSizeY int, tmpSize float64) (ok bool, totalHeight int) {
	maxWidth, maxHeight, tmpTotalHeight := 0, 0, 0
	for i := range lines {
		lineH, lineW := getTextBounds(lines[i], tmpSize)

		// only pad secondary lines, where its a multiline text
		if (i > 0) && len(lines) > 0 {
			lineH += linePadding
		}

		if lineW > maxWidth {
			maxWidth = lineW
		}
		if lineH > maxHeight {
			maxHeight = lineH
		}

		tmpTotalHeight += lineH
	}

	if (maxWidth > btnSizeX) || (tmpTotalHeight > btnSizeY) {
		return false, tmpTotalHeight
	}

	return true, tmpTotalHeight
}

func GetImageWithText(text string, textColor color.Color, backgroundColor color.Color, btnSize int) image.Image {
	return GetImageWithTextExt(text, textColor, backgroundColor, btnSize, ButtonOptions{})
}

func GetImageWithTextExt(text string, textColor color.Color, backgroundColor color.Color, btnSize int, opt ButtonOptions) image.Image {
	btnSizeX := btnSize - (2 * opt.MarginX)
	btnSizeY := btnSize - (2 * opt.MarginX)
	linepadding := 5
	maxsize := opt.MaxTextSize
	if maxsize <= 0 {
		maxsize = 60
	}

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	// split multiline strings
	lines := strings.Split(text, "\n")

	var size float64
	var totalHeight, prevTotalHeight int
	for size = 1; size < maxsize; size++ {
		var ok bool
		ok, totalHeight = testSizeWillFit(lines, linepadding, btnSizeX, btnSizeY, size)

		// when size exceed bounds, accept previous iteration
		if !ok {
			size -= 1
			totalHeight = prevTotalHeight
			break
		}
		prevTotalHeight = totalHeight
	}

	srcImg := image.NewUniform(textColor)
	dstImg := getSolidColourImage(backgroundColor, btnSize)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dstImg)
	c.SetSrc(srcImg)
	c.SetFontSize(size)
	c.SetClip(dstImg.Bounds())

	ypos := opt.MarginY
	if opt.AlignY == 0 {
		ypos = int((btnSizeY-totalHeight)/2) + opt.MarginY
	} else if opt.AlignY == 1 {
		ypos = int(btnSizeY - totalHeight)
		if ypos < opt.MarginY {
			ypos = opt.MarginY
		}
	}

	for i := range lines {
		lineh, linew := getTextBounds(lines[i], size)

		// only pad secondary lines, where its a multiline text
		if (i > 0) && len(lines) > 0 {
			lineh += linepadding
		}

		x := opt.MarginX
		if opt.AlignX == 0 {
			x = int((btnSizeX-linew)/2) + opt.MarginX
		} else if opt.AlignX == 1 {
			x = int(btnSizeX-linew) + opt.MarginX
		}

		pt := freetype.Pt(x, ypos+lineh)
		c.DrawString(lines[i], pt)
		ypos += lineh
	}

	return dstImg
}

func getTextWidth(text string, size float64) int {
	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(myfont, &truetype.Options{Size: size})
	w := 0
	for _, x := range text {
		awidth, _ := face.GlyphAdvance(rune(x))
		iwidthf := int(float64(awidth) / 64)
		w += iwidthf
	}

	return w
}

func getTextHeight(text string, size float64) int {
	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(myfont, &truetype.Options{Size: size})
	h := 0
	for _, y := range text {
		bounds, _, _ := face.GlyphBounds(rune(y))

		gheight := bounds.Max.Y - bounds.Min.Y
		iheightf := int(float64(gheight) / 64)
		if iheightf > h {
			h = iheightf
		}
	}

	return h
}

func getTextBounds(text string, size float64) (int, int) {
	return getTextHeight(text, size), getTextWidth(text, size)
}
