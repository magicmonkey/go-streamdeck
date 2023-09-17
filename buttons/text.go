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

// TextButton represents a button with text on it
type TextButton struct {
	label            string
	textColour       color.Color
	backgroundColour color.Color
	updateHandler    func(streamdeck.Button)
	btnIndex         int
	actionHandler    streamdeck.ButtonActionHandler
}

// GetImageForButton is the interface implemention to get the button's image as an image.Image
func (btn *TextButton) GetImageForButton(btnSize int) image.Image {
	img := getImageWithText(btn.label, btn.textColour, btn.backgroundColour, btnSize)
	return img
}

// SetButtonIndex is the interface implemention to set which button on the Streamdeck this is
func (btn *TextButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

// GetButtonIndex is the interface implemention to get which button on the Streamdeck this is
func (btn *TextButton) GetButtonIndex() int {
	return btn.btnIndex
}

// SetText allows the text on the button to be changed on the fly
func (btn *TextButton) SetText(label string) {
	btn.label = label
	btn.updateHandler(btn)
}

// SetTextColour allows the colour of the text on the button to be changed on the fly
func (btn *TextButton) SetTextColour(textColour color.Color) {
	btn.textColour = textColour
	btn.updateHandler(btn)
}

// SetBackgroundColor allows the background colour on the button to be changed on the fly
func (btn *TextButton) SetBackgroundColor(backgroundColour color.Color) {
	btn.backgroundColour = backgroundColour
	btn.updateHandler(btn)
}

// RegisterUpdateHandler is the interface implemention to let the engine give this button a callback to
// use to request that the button image is updated on the Streamdeck.
func (btn *TextButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

// SetActionHandler allows a ButtonActionHandler implementation to be
// set on this button, so that something can happen when the button is pressed.
func (btn *TextButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

// Pressed is the interface implementation for letting the engine notify that the button has been
// pressed.  This hands-off to the specified ButtonActionHandler if it has been set.
func (btn *TextButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

// NewTextButton creates a new TextButton with the specified text on it, in white on a black
// background.  The text will be set on a single line, and auto-sized to fill the button as best
// as possible.
func NewTextButton(label string) *TextButton {
	btn := NewTextButtonWithColours(label, color.White, color.Black)
	return btn
}

// NewTextButtonWithColours creates a new TextButton with the specified text on it, in the specified
// text and background colours.  The text will be set on a single line, and auto-sized to fill the
// button as best as possible.
func NewTextButtonWithColours(label string, textColour color.Color, backgroundColour color.Color) *TextButton {
	btn := &TextButton{label: label, textColour: textColour, backgroundColour: backgroundColour}
	return btn
}

func getImageWithText(text string, textColour color.Color, backgroundColour color.Color, btnSize int) image.Image {

	size := float64(18)

	myfont, err := truetype.Parse(gomedium.TTF)
	if err != nil {
		panic(err)
	}

	width := 0
	for size = 1; size < 60; size++ {
		width = getTextWidth(text, size)
		if width > btnSize {
			size = size - 1
			break
		}
	}

	srcImg := image.NewUniform(textColour)

	dstImg := image.NewRGBA(image.Rect(0, 0, btnSize, btnSize))
	draw.Draw(dstImg, dstImg.Bounds(), image.NewUniform(backgroundColour), image.Point{0, 0}, draw.Src)

	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dstImg)
	c.SetSrc(srcImg)
	c.SetFontSize(size)
	c.SetClip(dstImg.Bounds())

	x := int((btnSize - width) / 2) // Horizontally centre text
	y := int(btnSize / 2) + int(size / 3)  // Fudged vertical centre, erm, very "heuristic"

	pt := freetype.Pt(x, y)
	c.DrawString(text, pt)

	/*
		textWidth := 7 * len(text)
		fmt.Println(textWidth)

		f := &font.Drawer{
			Dst:  dstImg,
			Src:  src_img,
			Face: basicfont.Face7x13,
			Dot:  fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)},
		}
		f.DrawString(text)
	*/
	return dstImg
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
