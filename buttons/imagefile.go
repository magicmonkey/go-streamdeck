package buttons

import (
	"image"
	"os"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type ImageFileButton struct {
	filePath      string
	img           image.Image
	updateHandler func(streamdeck.Button)
	btnIndex      int
	actionHandler streamdeck.ButtonActionHandler
}

func (btn *ImageFileButton) GetImageForButton() image.Image {
	return btn.img
}

func (btn *ImageFileButton) SetButtonIndex(btnIndex int) {
	btn.btnIndex = btnIndex
}

func (btn *ImageFileButton) GetButtonIndex() int {
	return btn.btnIndex
}

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
	if err != nil {
		return err
	}
	btn.img = img
	return nil
}

func (btn *ImageFileButton) RegisterUpdateHandler(f func(streamdeck.Button)) {
	btn.updateHandler = f
}

func (btn *ImageFileButton) SetActionHandler(a streamdeck.ButtonActionHandler) {
	btn.actionHandler = a
}

func (btn *ImageFileButton) Pressed() {
	if btn.actionHandler != nil {
		btn.actionHandler.Pressed(btn)
	}
}

func NewImageFileButton(filePath string) (*ImageFileButton, error) {
	btn := &ImageFileButton{filePath: filePath}
	err := btn.loadImage()
	if err != nil {
		return nil, err
	}
	return btn, nil
}
