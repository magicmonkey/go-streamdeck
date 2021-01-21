package devices

import "errors"

func resetPacket(device string) ([]byte, error) {
	var size int
	var header []byte
	
	switch device {
	case "Streamdeck XL", "Streamdeck (original v2)":
		size = 32
		header = []byte{0x03, 0x02}
	case "Streamdeck Mini":
		size = 17
		header = []byte{0x0b, 0x63}
	default:
		return nil, errors.New("Device not supported!")
	}
	
	b := make([]byte, size)
	b[0] = header[0]
	b[1] = header[1]
	return b, nil
}

func brightnessPacket(device string) ([]byte, error) {
	var size int
	var header []byte

	switch device {
	case "Streamdeck XL", "Streamdeck (original v2)":
		size = 32
		header = []byte{0x03, 0x08}
	case "Streamdeck Mini":
		size = 17
		header = []byte{0x05, 0x55, 0xaa, 0xd1, 0x01}
	default:
		return nil, errors.New("Device is not supported!")
	}

	b := make([]byte, size)
	b[0] = header[0]
	b[1] = header[1]
	if device == "Streamdeck Mini" {
		b[2] = header[2]
		b[3] = header[3]
		b[4] = header[4]
	}
	return b, nil
}

