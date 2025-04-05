// Package clr provides efficient 16-bit and 32-bit color representation types and conversion between hexadecimal, RGB, RGBA and CSS named-colors.
// Along with W3C named-color hex value constants.
package clr

import (
	"fmt"
	"github.com/speedyhoon/cnst/clrn"
	"strconv"
	"strings"
)

// C16 represents a 16-bit color with 4 channels (red, green, blue and alpha) using 4-bits each (0-15 or 0-F).
type C16 uint16

const (
	x11             = 0x11     // Multiply to convert each Hex3 digit to a Hex6 digit.
	xF              = 0xF      // Isolate each channel using an AND operator.
	r, g, b         = 12, 8, 4 // Bitshift for each channel
	hexPrefix       = "#"
	c16AlphaDefault = "f"
)

// New16 returns a new 16-bit color where R, G, B, and A are converted from 0-255 range to 0-F hexadecimal.
// Use New16Hex for 4-bit hexadecimal integers ranging from 0 to F.
//
//	New16Hex(255, 102, 0, 255)
func New16(R, G, B, A uint8) C16 {
	return Round(C32(R)<<r) | Round(C32(G)<<g) | Round(C32(B)<<b) | Round(C32(A))
}

// New16Hex returns a new 16-bit color where R, G, B, and A are expected to be a hexadecimal ranging from 0 to F.
// Use New16 for 8-bit integers ranging from 0 to 255.
//
//	New16Hex(0xF, 0x6, 0x0, 0xF)
func New16Hex(R, G, B, A uint8) C16 {
	return C16(R)<<r | C16(G)<<g | C16(B)<<b | C16(A)
}

// New16Str parses string color and if valid returns a 16-bit color.
// Accepts any 3, 4, 6 or 8 digit hexadecimal, CSS named-color, RGB or RGBA color string.
//
//	FromHex4("rgba(255,102,0,255)")
func New16Str(color string) (_ C16, err error) {
	switch {
	case IsValidHex3(color):
		return FromHex3(color)
	case IsValidHex4(color):
		return FromHex4(color)

	case IsValidHex6(color):
		var c C32
		c, err = FromHex6(color)
		if err == nil {
			return c.To16(), nil
		}
	case IsValidHex8(color):
		var c C32
		c, err = FromHex8(color)
		if err == nil {
			return c.To16(), nil
		}
	case IsValidRGB(color):
		var c C32
		c, err = FromRGB(color)
		if err == nil {
			return c.To16(), nil
		}
	case IsValidRGBA(color):
		var c C32
		c, err = FromRGBA(color)
		if err == nil {
			return c.To16(), nil
		}
	default:
		if c := FromName(color); c != 0 {
			return c.To16(), nil
		}
	}

	return 0, ErrInvalid
}

// FromHex3 converts a 3-digit hexadecimal string into a 16-bit color.
//
//	FromHex3("#F60")
func FromHex3(c string) (C16, error) {
	return FromHex4(c + c16AlphaDefault)
}

// FromHex4 converts a 4-digit hexadecimal string into a 16-bit color.
//
//	FromHex4("#F60F")
func FromHex4(color string) (C16, error) {
	color = strings.TrimPrefix(color, hexPrefix)
	u64, err := strconv.ParseUint(color, 16, 16)
	if err != nil {
		return 0, ErrInvalid
	}

	return C16(u64), nil
}

// To32 converts a 16-bit color to a 32-bit color.
func (c C16) To32() C32 {
	return C32(c)&0xF000*x11<<r | C32(c)&0xF00*x11<<g | C32(c)&0xF0*x11<<b | C32(c)&xF*x11
}

// String returns the shortest possible string representation of color without loosing any color accuracy.
// Returns:
//
//	"red" or
//	a 4-digit hexadecimal color string when the alpha channel is <= 254 (like `#f608`) otherwise,
//	a 3-digit hexadecimal color string when the alpha channel equals 255 (like `#f60`).
func (c C16) String() string {
	if c&xF != xF {
		return c.Hex4()
	}
	if c == 0xF00F {
		return clrn.Red
	}
	return c.Hex3()
}

// Lossy returns the shortest possible string representation of color ignoring the alpha channel.
// Returns:
//
//	"red" otherwise,
//	a 3-digit hexadecimal color string (like `#f60`).
func (c C16) Lossy() string {
	if c == 0xF00F {
		return clrn.Red
	}
	return c.Hex3()
}

// Hex3 returns a 3-digit hexadecimal color string without the alpha channel like `#f00`.
func (c C16) Hex3() string {
	return fmt.Sprintf("#%03x", uint16(c>>b))
}

// Hex4 returns a 4-digit hexadecimal color string with the alpha channel like `#f00f`.
func (c C16) Hex4() string {
	return fmt.Sprintf("#%04x", uint16(c))
}

// Hex6 returns a 6-digit hexadecimal color string with the alpha channel like `#ff0000`.
func (c C16) Hex6() string {
	return c.To32().Hex6()
}

// Hex8 returns an 8-digit hexadecimal color string with the alpha channel like `#ff0000ff`.
func (c C16) Hex8() string {
	return c.To32().Hex8()
}

// RGB returns an RGB color string without the alpha channel like `rgb(255,0,0)`.
func (c C16) RGB() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c>>r*x11, c>>g&xF*x11, c>>b&xF*x11)
}

// RGBA returns an RGBA color string with the alpha channel like `rgba(255,0,0,255)`.
func (c C16) RGBA() string {
	return fmt.Sprintf("rgb(%d,%d,%d,%d)", c>>r*x11, c>>g&xF*x11, c>>b&xF*x11, c&xF*x11)
}
