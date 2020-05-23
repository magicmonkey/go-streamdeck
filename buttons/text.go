package buttons

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font/gofont/gomedium"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type TextButton struct {
	label            string
	textColour       color.Color
	backgroundColour color.Color
	updateHandler    func(streamdeck.Button)
	btnIndex         int
	actionHandler    streamdeck.ButtonActionHandler
}

func (btn *TextButton) GetImageForButton() image.Image {
	img := getImageWithText(btn.label, btn.textColour, btn.backgroundColour)
	return img
}

func (btn *TextButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

func (btn *TextButton) GetButtonIndex() int {
	return btn.btnIndex
}

func (btn *TextButton) SetText(label string) {
	btn.label = label
	btn.updateHandler(btn)
}

func (btn *TextButton) SetTextColour(textColour color.Color) {
	btn.textColour = textColour
	btn.updateHandler(btn)
}

func (btn *TextButton) SetBackgroundColor(backgroundColour color.Color) {
	btn.backgroundColour = backgroundColour
	btn.updateHandler(btn)
}

func (btn *TextButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

func (btn *TextButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

func (btn *TextButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

func NewTextButton(label string) *TextButton {
	btn := NewTextButtonWithColours(label, color.White, color.Black)
	return btn
}

func NewTextButtonWithColours(label string, textColour color.Color, backgroundColour color.Color) *TextButton {
	btn := &TextButton{label: label, textColour: textColour, backgroundColour: backgroundColour}
	return btn
}

func getImageWithText(text string, textColour color.Color, backgroundColour color.Color) image.Image {

	ButtonSize := 96
	size := float64(18)

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	width := 0
	for size = 1; size < 60; size++ {
		width = getTextWidth(text, size)
		if width > 90 {
			size = size - 1
			break
		}
	}

	src_img := image.NewUniform(textColour)

	dst_img := image.NewRGBA(image.Rect(0, 0, ButtonSize, ButtonSize))
	draw.Draw(dst_img, dst_img.Bounds(), image.NewUniform(backgroundColour), image.Point{0, 0}, draw.Src)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dst_img)
	c.SetSrc(src_img)
	c.SetFontSize(size)
	c.SetClip(dst_img.Bounds())

	x := int((96 - width) / 2) // Horizontally centre text
	y := int(50 + (size / 3))  // Fudged vertical centre, erm, very "heuristic"

	pt := freetype.Pt(x, y)
	c.DrawString(text, pt)

	/*
		textWidth := 7 * len(text)
		fmt.Println(textWidth)

		f := &font.Drawer{
			Dst:  dst_img,
			Src:  src_img,
			Face: basicfont.Face7x13,
			Dot:  fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)},
		}
		f.DrawString(text)
	*/
	return dst_img
}

func getTextWidth(text string, size float64) int {

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	// Calculate width of string
	width := 0
	face := truetype.NewFace(myfont, &truetype.Options{Size: size})
	for _, x := range text {
		awidth, _ := face.GlyphAdvance(rune(x))
		iwidthf := int(float64(awidth) / 64)
		width += iwidthf
	}

	return width
}
