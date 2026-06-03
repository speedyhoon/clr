package clr

// ToBGRA32 uses 32 bits per pixel with 8 bits per component
// with the order Blue, Green, Red, and Alpha.
func (c C32) ToBGRA32() []byte {
	return []byte{byte(c >> 8), byte(c >> 16), byte(c >> 24), byte(c)}
}

// ToBGR032 packs a pixel into 32-bit BGR format (B, G, R, 8-bit padding),
// where the padding byte is always set to 0x00.
func (c C32) ToBGR032() []byte {
	return []byte{byte(c >> 8), byte(c >> 16), byte(c >> 24), 0}
}

// ToBGRX32 packs a pixel into 32-bit BGRX format (B, G, R, 8-bit padding),
// where the padding byte is always set to 0xFF.
func (c C32) ToBGRX32() []byte {
	return []byte{byte(c >> 8), byte(c >> 16), byte(c >> 24), 255}
}
