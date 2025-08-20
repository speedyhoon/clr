package clr

import "math/rand"

// RandomC16 returns a random 16-bit color from 27 possible web-safe primary colors with 100% opacity.
// Each red, green and blue channel uses either 0, 7 or F as their hexadecimal digits.
func RandomC16() C16 {
	const size = 3
	c := [size]C16{0, 0x7, xF} // Options to randomly select shades from.
	return c[rand.Intn(size)]<<r | c[rand.Intn(size)]<<g | c[rand.Intn(size)]<<b | xF
}

// RandomC32 returns a random 32-bit color with 100% opacity.
func RandomC32() C32 {
	const MaxUint24 = 1<<24 - 1
	return C32(rand.Intn(MaxUint24))<<g | xFF
}
