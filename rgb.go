package clr

import (
	"fmt"
	"strconv"
)

func (c C32) RGB() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c>>rr, c>>gg&xFF, c>>bb&xFF)
}

func (c C32) RGBA() string {
	return fmt.Sprintf("rgb(%d,%d,%d,%d)", c>>rr, c>>gg&xFF, c>>bb&xFF, c&xFF)
}

const rgbChannels, rgbaChannels int8 = 3, 4

// FromRGB converts 3 channels (red, green, and blue) ranging from 0 to 255 in base 10 into C32.
// All formats are supported, provided a non-numeric character separates each value. For example:
//
//	rgb(0,0,0)
//	0_0 0
func FromRGB(color string) (h C32, err error) {
	var R, G, B uint8
	R, G, B, _, err = parseRGBA(color, rgbChannels)
	return New32(R, G, B, xFF), err
}

// FromRGBA converts 4 channels (red, green, blue, and alpha) ranging from 0 to 255 in base 10 into C32.
// All formats are supported, provided a non-numeric character separates each value. For example:
//
//	rgba(0, 0, 0, 0)
//	0!0-0=0
func FromRGBA(color string) (h C32, err error) {
	var R, G, B, A uint8
	R, G, B, A, err = parseRGBA(color, rgbaChannels)
	return New32(R, G, B, A), err
}

func parseRGBA(color string, channelsQty int8) (r, g, b, a uint8, err error) {
	var index int8
	var prevIsNumeric bool
	ptr := []*uint8{&r, &g, &b, &a}
	var start, i int

	var c rune
	for i, c = range color {
		if isNumeric(c) {
			if !prevIsNumeric {
				start = i
				prevIsNumeric = true
			}

			continue
		}

		if prevIsNumeric {
			*ptr[index], err = parse8bitChannel(color[start:i])

			index++
			if index >= channelsQty || err != nil {
				return
			}

			prevIsNumeric = false
		}
	}

	if index+1 != channelsQty {
		return r, g, b, a, ErrNot3Channels
	}

	*ptr[index], err = parse8bitChannel(color[start:i])
	return
}

// parse8bitChannel converts a base 10 decimal string into an uint8.
// The expected range is 0-255 and is used to convert the RGB format.
func parse8bitChannel(c string) (uint8, error) {
	C, err := strconv.ParseUint(c, 10, 8)
	if err != nil {
		return 0, ErrExceedRange
	}
	return uint8(C), nil
}

func isNumeric(c rune) bool {
	return c >= '0' && c <= '9'
}
