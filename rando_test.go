package clr_test

import (
	"github.com/speedyhoon/clr"
	"testing"
)

func TestRandomC16(t *testing.T) {
	colorList := map[clr.C16]uint8{
		0x000f: 0,
		0x007f: 0,
		0x00ff: 0,
		0x070f: 0,
		0x077f: 0,
		0x07ff: 0,
		0x0f0f: 0,
		0x0f7f: 0,
		0x0fff: 0,
		0x700f: 0,
		0x707f: 0,
		0x70ff: 0,
		0x770f: 0,
		0x777f: 0,
		0x77ff: 0,
		0x7f0f: 0,
		0x7f7f: 0,
		0x7fff: 0,
		0xf00f: 0,
		0xf07f: 0,
		0xf0ff: 0,
		0xf70f: 0,
		0xf77f: 0,
		0xf7ff: 0,
		0xff0f: 0,
		0xff7f: 0,
		0xffff: 0,
	}
	for i := 0; i < 999; i++ {
		c := clr.RandomC16()
		if _, ok := colorList[c]; !ok {
			t.Errorf("RandomC16() returned an unexpected color %s", c.Hex4())
			return
		}

		colorList[c]++
	}
	for c, qty := range colorList {
		if qty == 0 {
			t.Errorf("RandomC16() missed expected color %s", c.Hex4())
		}
	}
}

func TestRandomC32(t *testing.T) {
	var prev clr.C32
	for i := 0; i < 10; i++ {
		c := clr.RandomC32()

		if c&0xff != 0xff {
			t.Errorf("RandomC32() returned unexpected opacity value %s", c.Hex8())
		}

		if prev == c {
			t.Errorf("RandomC32() returned the same value %s as previous %s", c.Hex8(), prev.Hex8())
		}
		prev = c
	}
}
