package clr

import (
	"fmt"
	"github.com/speedyhoon/cnst/clrn"
	"strconv"
	"strings"
)

// C32 represents a 32-bit color with four channels (red, green, blue and alpha) using 8-bits each (0-255 or 0-FF).
type C32 uint32

const (
	xFF                = 0xFF      // Isolate each channel using an AND operator.
	f11        float32 = x11       // Used for rounding.
	rr, gg, bb         = 24, 16, 8 // Bitshift for each channel
	bitSize            = 32
)

func New32RGBA(R, G, B, A uint8) C32 {
	return C32(R)<<rr | C32(G)<<gg | C32(B)<<bb | C32(A)
}

// New32Str parses string color and if valid returns a 32-bit color.
func New32Str(color string) (_ C32, err error) {
	switch {
	case IsValidHex3(color):
		var c C16
		c, err = FromHex3(color)
		if err == nil {
			return c.To32(), nil
		}
	case IsValidHex4(color):
		var c C16
		c, err = FromHex4(color)
		if err == nil {
			return c.To32(), nil
		}
	case IsValidHex6(color):
		return FromHex6(color)
	case IsValidHex8(color):
		return FromHex8(color)
	case IsValidRGB(color):
		return FromRGB(color)
	case IsValidRGBA(color):
		return FromRGBA(color)
	default:
		if c := FromName(color); c != 0 {
			return c, nil
		}
	}

	return 0, ErrInvalid
}

// FromHex6 converts a 6-digit hexadecimal string into a 32-bit color.
//
//	FromHex6("#FF6600")
func FromHex6(color string) (C32, error) {
	if !IsValidHex6(color) {
		return 0, ErrInvalid
	}

	color = strings.TrimPrefix(color, hexPrefix)
	u64, err := strconv.ParseUint(color, gg, rr)
	if err != nil {
		return 0, ErrInvalid
	}

	return C32(u64)<<bb | xFF, nil
}

// FromHex8 converts an 8-digit hexadecimal string into a 32-bit color.
//
//	FromHex8("#FF6600FF")
func FromHex8(color string) (C32, error) {
	color = strings.TrimPrefix(color, hexPrefix)
	u64, err := strconv.ParseUint(color, gg, bitSize)
	if err != nil {
		return 0, ErrInvalid
	}

	return C32(u64), nil
}

// To16 is a lossy conversion to a 16-bit color.
func (c C32) To16() C16 {
	return Round(c&0xFF000000>>rr)<<r | Round(c&0xFF0000>>gg)<<g | Round(c&0xFF00>>bb)<<b | Round(c&xFF)
}

// Round converts a single C32 channel into a single C16 channel.
// The returned C16 channel is rounded to the closest hexadecimal integer
// for optimal color accuracy.
func Round(channel C32) C16 {
	// Adding 0.5 is more performant than calling math.Round().
	return C16(float32(channel)/f11 + 0.5)
}

// Hex3 returns a lossy 3-digit hexadecimal color string without the alpha channel like `#f00`.
func (c C32) Hex3() string {
	return c.To16().Hex3()
}

// Hex4 returns a lossy 4-digit hexadecimal color string with the alpha channel like `#f00f`.
func (c C32) Hex4() string {
	return c.To16().Hex4()
}

// Hex6 returns a 6-digit hexadecimal color string with the alpha channel like `#ff0000`.
func (c C32) Hex6() string {
	return fmt.Sprintf("#%06x", uint32(c>>bb))
}

// Hex8 returns an 8-digit hexadecimal color string with the alpha channel like `#ff0000ff`.
func (c C32) Hex8() string {
	return fmt.Sprintf("#%08x", uint32(c))
}

// String returns the shortest possible string representation of color without loosing any color accuracy.
func (c C32) String() string {
	// Hex 3 or 6 digit.
	if c&xFF == xFF {
		if c == Red {
			return clrn.Red
		}

		if c.Is24bit() {
			// Check if the color is a named-color that's shorter than the 6-digit hexadecimal string.
			switch c {
			case Azure:
				return clrn.Azure
			case Beige:
				return clrn.Beige
			case Bisque:
				return clrn.Bisque
			case Brown:
				return clrn.Brown
			case Coral:
				return clrn.Coral
			case Gold:
				return clrn.Gold
			case Gray:
				return clrn.Gray
			case Green:
				return clrn.Green
			case Indigo:
				return clrn.Indigo
			case Ivory:
				return clrn.Ivory
			case Khaki:
				return clrn.Khaki
			case Linen:
				return clrn.Linen
			case Maroon:
				return clrn.Maroon
			case Navy:
				return clrn.Navy
			case Olive:
				return clrn.Olive
			case Orange:
				return clrn.Orange
			case Orchid:
				return clrn.Orchid
			case Peru:
				return clrn.Peru
			case Pink:
				return clrn.Pink
			case Plum:
				return clrn.Plum
			case Purple:
				return clrn.Purple
			case Salmon:
				return clrn.Salmon
			case Sienna:
				return clrn.Sienna
			case Silver:
				return clrn.Silver
			case Snow:
				return clrn.Snow
			case Tan:
				return clrn.Tan
			case Teal:
				return clrn.Teal
			case Tomato:
				return clrn.Tomato
			case Violet:
				return clrn.Violet
			case Wheat:
				return clrn.Wheat
			}
			return c.Hex6()
		}

		return c.Hex3()
	} else if c.Is24bit() || c&xFF%x11 != 0 {
		return c.Hex8()
	}

	return c.Hex4()
}

// Is24bit determines if a color's red, green and blue channels can be represented as a 16-bit color without
// any color accuracy loss. The alpha (transparency) channel is ignored.
func (c C32) Is24bit() bool {
	return c&0xFF000000%x11 != 0 || c&0xFF0000%x11 != 0 || c&0xFF00%x11 != 0
}

// Is16bit determines if a C32 color can be converted to C16 without any accuracy loss.
func (c C32) Is16bit() bool {
	return c&0xFF000000%x11 == 0 && c&0xFF0000%x11 == 0 && c&0xFF00%x11 == 0 && c&0xFF%x11 == 0
}
