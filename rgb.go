package clr

import (
	"fmt"
	"strconv"
	"strings"
)

func (c C32) RGB() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c>>rr, c>>gg&xFF, c>>bb&xFF)
}

func (c C32) RGBA() string {
	return fmt.Sprintf("rgb(%d,%d,%d,%d)", c>>rr, c>>gg&xFF, c>>bb&xFF, c&xFF)
}

// FromRGB converts 3 (RGB) or 4 (RGBA) comma separated channels ranging from 0 to 255 in base 10 into C32.
// Expected color string formats are:
//
//	rgb(0,0,0)
//	rgba(0, 0, 0, 0)
//	0,0,0
//	0,0,0,0
func FromRGB(color string) (h C32, err error) {
	color = TrimRGBPrefix(color)

	channels := strings.Split(color, ",")
	if len(channels) != 3 {
		return 0, ErrNot3Channels
	}

	var R, G, B uint8
	R, err = parse8bitChannel(channels[0])
	if err != nil {
		return
	}

	G, err = parse8bitChannel(channels[1])
	if err != nil {
		return
	}

	B, err = parse8bitChannel(channels[2])
	if err != nil {
		return
	}

	return New32(R, G, B, xFF), nil
}

func FromRGBA(color string) (h C32, err error) {
	color = TrimRGBPrefix(color)

	channels := strings.Split(color, ",")
	if len(channels) != 4 {
		return 0, ErrInvalid
	}

	var R, G, B, A uint8

	R, err = parse8bitChannel(channels[0])
	if err != nil {
		return
	}

	G, err = parse8bitChannel(channels[1])
	if err != nil {
		return
	}

	B, err = parse8bitChannel(channels[2])
	if err != nil {
		return
	}

	A, err = parse8bitChannel(channels[3])
	if err != nil {
		return
	}

	return New32(R, G, B, A), nil
}

func TrimRGBPrefix(color string) string {
	color = strings.ReplaceAll(color, " ", "")
	color = strings.ReplaceAll(color, "\t", "")
	color = strings.TrimRight(color, ")")
	color = strings.ToLower(color)
	color = strings.TrimPrefix(color, "rgb(")
	return strings.TrimPrefix(color, "rgba(")
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
