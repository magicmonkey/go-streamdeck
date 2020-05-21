package streamdeck

type TextButton struct {
	label         string
	updateHandler func(Button)
	btnIndex      int
	actionHandler ButtonActionHandler
}

func (tb *TextButton) GetButtonType() string {
	return "text"
}

func (tb *TextButton) GetImageFile() string {
	return ""
}

func (tb *TextButton) GetText() string {
	return tb.label
}

func (tb *TextButton) SetButtonIndex(btnIndex int) {
	tb.btnIndex = btnIndex
}

func (tb *TextButton) GetButtonIndex() int {
	return tb.btnIndex
}

func (tb *TextButton) SetText(label string) {
	tb.label = label
	tb.updateHandler(tb)
}

func (tb *TextButton) RegisterUpdateHandler(f func(Button)) {
	tb.updateHandler = f
}

func (tb *TextButton) SetActionHandler(a ButtonActionHandler) {
	a.SetButton(tb)
	tb.actionHandler = a
}

func (tb *TextButton) Pressed() {
	if tb.actionHandler != nil {
		tb.actionHandler.Pressed()
	}
}

func NewTextButton(label string) *TextButton {
	tb := &TextButton{label: label}
	return tb
}
