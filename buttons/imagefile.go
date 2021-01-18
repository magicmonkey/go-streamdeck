package buttons

import (
	"github.com/disintegration/gift"
	"image"
	"image/draw"
	"os"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

// ImageFileButton represents a button with an image on it, where the image is loaded
// from a file.
type ImageFileButton struct {
	filePath      string
	img           image.Image
	updateHandler func(streamdeck.Button)
	btnIndex      int
	actionHandler streamdeck.ButtonActionHandler
}

// GetImageForButton is the interface implemention to get the button's image as an image.Image
func (btn *ImageFileButton) GetImageForButton(btnSize int) image.Image {
	// Resize the image to what the button wants
	g := gift.New(gift.Resize(btnSize, btnSize, gift.LanczosResampling))
	newimg := image.NewRGBA(image.Rect(0, 0, btn.img.Bounds().Max.X, btn.img.Bounds().Max.Y))
	g.Draw(newimg, btn.img)
	return newimg
}

// SetButtonIndex is the interface implemention to set which button on the Streamdeck this is
func (btn *ImageFileButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

// GetButtonIndex is the interface implemention to get which button on the Streamdeck this is
func (btn *ImageFileButton) GetButtonIndex() int {
	return btn.btnIndex
}

// SetFilePath allows the image file to be changed on the fly
func (btn *ImageFileButton) SetFilePath(filePath string) error {
	btn.filePath = filePath
	err := btn.loadImage()
	if err != nil {
		return err
	}
	btn.updateHandler(btn)
	return nil
}

func (btn *ImageFileButton) loadImage() error {
	f, err := os.Open(btn.filePath)
	if err != nil {
		return err
	}
	img, _, err := image.Decode(f)

	// We want the image as an RGBA, so convert it if it isn't
	var newimg *image.RGBA
	newimg, ok := img.(*image.RGBA)
	if !ok {
		newimg = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
		draw.Draw(newimg, newimg.Bounds(), img, image.Point{0, 0}, draw.Src)
	}

	if err != nil {
		return err
	}
	btn.img = newimg
	return nil
}

// RegisterUpdateHandler is the interface implemention to let the engine give this button a callback to
// use to request that the button image is updated on the Streamdeck.
func (btn *ImageFileButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

// SetActionHandler allows a ButtonActionHandler implementation to be
// set on this button, so that something can happen when the button is pressed.
func (btn *ImageFileButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

// Pressed is the interface implementation for letting the engine notify that the button has been
// pressed.  This hands-off to the specified ButtonActionHandler if it has been set.
func (btn *ImageFileButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

// NewImageFileButton creates a new ImageFileButton with the specified image on it
func NewImageFileButton(filePath string) (*ImageFileButton, error) {
	btn := &ImageFileButton{filePath: filePath}
	err := btn.loadImage()
	if err != nil {
		return nil, err
	}
	return btn, nil
}
