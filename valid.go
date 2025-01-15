package clr

import (
	"errors"
	"regexp"
)

var (
	ErrInvalid     = errors.New("invalid value: expected a color name or 3, 4, 6, or 8 hexadecimal digits")
	ErrExceedRange = errors.New("expected an integer between 0 and 255")

	Hex3 = regexp.MustCompile(`^#?[[:xdigit:]]{3}$`)
	Hex4 = regexp.MustCompile(`^#?[[:xdigit:]]{4}$`)
	Hex6 = regexp.MustCompile(`^#?[[:xdigit:]]{6}$`)
	Hex8 = regexp.MustCompile(`^#?[[:xdigit:]]{8}$`)
	RGB  = regexp.MustCompile(`^rgb\(\d{1,3},\d{1,3},\d{1,3}\)$`)
	RGBA = regexp.MustCompile(`^rgba\(\d{1,3},\d{1,3},\d{1,3},\d{1,3}\)$`)
)

func IsValidHex3(c string) bool {
	return Hex3.MatchString(c)
}

func IsValidHex4(c string) bool {
	return Hex4.MatchString(c)
}

func IsValidHex6(c string) bool {
	return Hex6.MatchString(c)
}

func IsValidHex8(c string) bool {
	return Hex8.MatchString(c)
}

func IsValidRGB(c string) bool {
	return RGB.MatchString(c)
}

func IsValidRGBA(c string) bool {
	return RGBA.MatchString(c)
}
