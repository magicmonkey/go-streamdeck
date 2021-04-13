package devices

// resetPacket17 gives the reset packet for devices which need it to be 17 bytes long
func resetPacket17() []byte {
	pkt := make([]byte, 17)
	pkt[0] = 0x0b
	pkt[1] = 0x63
	return pkt
}

// resetPacket32 gives the reset packet for devices which need it to be 32 bytes long
func resetPacket32() []byte {
	pkt := make([]byte, 32)
	pkt[0] = 0x03
	pkt[1] = 0x02
	return pkt
}

// brightnessPacket17 gives the brightness packet for devices which need it to be 17 bytes long
func brightnessPacket17() []byte {
	pkt := make([]byte, 5)
	pkt[0] = 0x05
	pkt[1] = 0x55
	pkt[2] = 0xaa
	pkt[3] = 0xd1
	pkt[4] = 0x01
	return pkt
}

// brightnessPacket32 gives the brightness packet for devices which need it to be 32 bytes long
func brightnessPacket32() []byte {
	pkt := make([]byte, 2)
	pkt[0] = 0x03
	pkt[1] = 0x08
	return pkt
}
