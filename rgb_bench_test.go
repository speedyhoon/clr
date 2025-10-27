package clr_test

import (
	"github.com/speedyhoon/clr"
	"testing"
)

var color clr.C32

const input = " \t \nrgba(238, 212, 95, 255)\n\t "

func BenchmarkFromRGBA(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		color, err = clr.FromRGBA(input)
		if err != nil {
			panic(err)
		}
	}
}
