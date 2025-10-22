package clr_test

import (
	"fmt"
	"github.com/speedyhoon/clr"
	"testing"
)

func TestFromRGB(t *testing.T) {
	tests := []struct {
		name  string
		color string
		want  clr.C32
		err   error
	}{
		{color: "rgb(255, 255, 255)", want: clr.White},
		{color: "rgb(255,255,255)", want: clr.White},
		{color: "rgb(255 255 255)", want: clr.TransparentBlack, err: clr.ErrNot3Channels},
		{color: "rgb(0, 0, 0)", want: clr.Black},
		{color: "rgb(0,0,0)", want: clr.Black},
		{color: "rgb(0 0 0)", want: clr.TransparentBlack, err: clr.ErrNot3Channels},
		{color: "rgb(255,258,255)", want: clr.TransparentBlack, err: clr.ErrExceedRange},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]=`%v`", i, tt.color), func(t *testing.T) {
			gotColor, err := clr.FromRGB(tt.color)
			if err != tt.err {
				t.Errorf("FromRGB() got error: `%v`, want: `%v`", err, tt.err)
				return
			}
			if gotColor != tt.want {
				t.Errorf("FromRGB() got color = %v, want %v", gotColor.RGBA(), tt.want.RGBA())
			}
		})
	}
}
