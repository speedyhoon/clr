package clr

// ToBGRA32 uses 32 bits per pixel with 8 bits per component
// with the order Blue, Green, Red, and Alpha.
func (c C32) ToBGRA32() []byte {
	return []byte{byte(c >> 8), byte(c >> 16), byte(c >> 24), byte(c)}
}

// ToBGRX32 uses 32 bits per pixel with 8 bits per component
// with the order Blue, Green, Red, and a fourth padding byte when transparency is not required.
func (c C32) ToBGRX32() []byte {
	return []byte{byte(c >> 8), byte(c >> 16), byte(c >> 24), 0}
}
