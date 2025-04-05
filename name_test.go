package clr_test

import (
	"github.com/speedyhoon/clr"
	"github.com/speedyhoon/cnst/clrn"
	"testing"
)

func TestFromName_Hex6(t *testing.T) {
	names := map[string]string{
		clrn.AliceBlue:            "#f0f8ff",
		clrn.AntiqueWhite:         "#faebd7",
		clrn.Aqua:                 "#00ffff",
		clrn.Aquamarine:           "#7fffd4",
		clrn.Azure:                "#f0ffff",
		clrn.Beige:                "#f5f5dc",
		clrn.Bisque:               "#ffe4c4",
		clrn.Black:                "#000000",
		clrn.BlanchedAlmond:       "#ffebcd",
		clrn.Blue:                 "#0000ff",
		clrn.BlueViolet:           "#8a2be2",
		clrn.Brown:                "#a52a2a",
		clrn.BurlyWood:            "#deb887",
		clrn.CadetBlue:            "#5f9ea0",
		clrn.Chartreuse:           "#7fff00",
		clrn.Chocolate:            "#d2691e",
		clrn.Coral:                "#ff7f50",
		clrn.CornflowerBlue:       "#6495ed",
		clrn.Cornsilk:             "#fff8dc",
		clrn.Crimson:              "#dc143c",
		clrn.Cyan:                 "#00ffff",
		clrn.DarkBlue:             "#00008b",
		clrn.DarkCyan:             "#008b8b",
		clrn.DarkGoldenrod:        "#b8860b",
		clrn.DarkGray:             "#a9a9a9",
		clrn.DarkGreen:            "#006400",
		clrn.DarkGrey:             "#a9a9a9",
		clrn.DarkKhaki:            "#bdb76b",
		clrn.DarkMagenta:          "#8b008b",
		clrn.DarkOliveGreen:       "#556b2f",
		clrn.DarkOrange:           "#ff8c00",
		clrn.DarkOrchid:           "#9932cc",
		clrn.DarkRed:              "#8b0000",
		clrn.DarkSalmon:           "#e9967a",
		clrn.DarkSeaGreen:         "#8fbc8f",
		clrn.DarkSlateBlue:        "#483d8b",
		clrn.DarkSlateGray:        "#2f4f4f",
		clrn.DarkSlateGrey:        "#2f4f4f",
		clrn.DarkTurquoise:        "#00ced1",
		clrn.DarkViolet:           "#9400d3",
		clrn.DeepPink:             "#ff1493",
		clrn.DeepSkyBlue:          "#00bfff",
		clrn.DimGray:              "#696969",
		clrn.DimGrey:              "#696969",
		clrn.DodgerBlue:           "#1e90ff",
		clrn.Firebrick:            "#b22222",
		clrn.FloralWhite:          "#fffaf0",
		clrn.ForestGreen:          "#228b22",
		clrn.Fuchsia:              "#ff00ff",
		clrn.Gainsboro:            "#dcdcdc",
		clrn.GhostWhite:           "#f8f8ff",
		clrn.Gold:                 "#ffd700",
		clrn.Goldenrod:            "#daa520",
		clrn.Gray:                 "#808080",
		clrn.Green:                "#008000",
		clrn.GreenYellow:          "#adff2f",
		clrn.Grey:                 "#808080",
		clrn.Honeydew:             "#f0fff0",
		clrn.HotPink:              "#ff69b4",
		clrn.IndianRed:            "#cd5c5c",
		clrn.Indigo:               "#4b0082",
		clrn.Ivory:                "#fffff0",
		clrn.Khaki:                "#f0e68c",
		clrn.Lavender:             "#e6e6fa",
		clrn.LavenderBlush:        "#fff0f5",
		clrn.LawnGreen:            "#7cfc00",
		clrn.LemonChiffon:         "#fffacd",
		clrn.LightBlue:            "#add8e6",
		clrn.LightCoral:           "#f08080",
		clrn.LightCyan:            "#e0ffff",
		clrn.LightGoldenrodYellow: "#fafad2",
		clrn.LightGray:            "#d3d3d3",
		clrn.LightGreen:           "#90ee90",
		clrn.LightGrey:            "#d3d3d3",
		clrn.LightPink:            "#ffb6c1",
		clrn.LightSalmon:          "#ffa07a",
		clrn.LightSeaGreen:        "#20b2aa",
		clrn.LightSkyBlue:         "#87cefa",
		clrn.LightSlateGray:       "#778899",
		clrn.LightSlateGrey:       "#778899",
		clrn.LightSteelBlue:       "#b0c4de",
		clrn.LightYellow:          "#ffffe0",
		clrn.Lime:                 "#00ff00",
		clrn.LimeGreen:            "#32cd32",
		clrn.Linen:                "#faf0e6",
		clrn.Magenta:              "#ff00ff",
		clrn.Maroon:               "#800000",
		clrn.MediumAquamarine:     "#66cdaa",
		clrn.MediumBlue:           "#0000cd",
		clrn.MediumOrchid:         "#ba55d3",
		clrn.MediumPurple:         "#9370db",
		clrn.MediumSeaGreen:       "#3cb371",
		clrn.MediumSlateBlue:      "#7b68ee",
		clrn.MediumSpringGreen:    "#00fa9a",
		clrn.MediumTurquoise:      "#48d1cc",
		clrn.MediumVioletRed:      "#c71585",
		clrn.MidnightBlue:         "#191970",
		clrn.MintCream:            "#f5fffa",
		clrn.MistyRose:            "#ffe4e1",
		clrn.Moccasin:             "#ffe4b5",
		clrn.NavajoWhite:          "#ffdead",
		clrn.Navy:                 "#000080",
		clrn.OldLace:              "#fdf5e6",
		clrn.Olive:                "#808000",
		clrn.OliveDrab:            "#6b8e23",
		clrn.Orange:               "#ffa500",
		clrn.OrangeRed:            "#ff4500",
		clrn.Orchid:               "#da70d6",
		clrn.PaleGoldenrod:        "#eee8aa",
		clrn.PaleGreen:            "#98fb98",
		clrn.PaleTurquoise:        "#afeeee",
		clrn.PaleVioletRed:        "#db7093",
		clrn.PapayaWhip:           "#ffefd5",
		clrn.PeachPuff:            "#ffdab9",
		clrn.Peru:                 "#cd853f",
		clrn.Pink:                 "#ffc0cb",
		clrn.Plum:                 "#dda0dd",
		clrn.PowderBlue:           "#b0e0e6",
		clrn.Purple:               "#800080",
		clrn.RebeccaPurple:        "#663399",
		clrn.Red:                  "#ff0000",
		clrn.RosyBrown:            "#bc8f8f",
		clrn.RoyalBlue:            "#4169e1",
		clrn.SaddleBrown:          "#8b4513",
		clrn.Salmon:               "#fa8072",
		clrn.SandyBrown:           "#f4a460",
		clrn.SeaGreen:             "#2e8b57",
		clrn.Seashell:             "#fff5ee",
		clrn.Sienna:               "#a0522d",
		clrn.Silver:               "#c0c0c0",
		clrn.SkyBlue:              "#87ceeb",
		clrn.SlateBlue:            "#6a5acd",
		clrn.SlateGray:            "#708090",
		clrn.SlateGrey:            "#708090",
		clrn.Snow:                 "#fffafa",
		clrn.SpringGreen:          "#00ff7f",
		clrn.SteelBlue:            "#4682b4",
		clrn.Tan:                  "#d2b48c",
		clrn.Teal:                 "#008080",
		clrn.Thistle:              "#d8bfd8",
		clrn.Tomato:               "#ff6347",
		clrn.Turquoise:            "#40e0d0",
		clrn.Violet:               "#ee82ee",
		clrn.Wheat:                "#f5deb3",
		clrn.White:                "#ffffff",
		clrn.WhiteSmoke:           "#f5f5f5",
		clrn.Yellow:               "#ffff00",
		clrn.YellowGreen:          "#9acd32",
	}

	const expectedQty = 148
	if qty := len(names); qty != expectedQty {
		t.Errorf("have %d color names, expected %d\n", qty, expectedQty)
	}

	for name, expected := range names {
		hex6 := clr.FromName(name).Hex6()
		if hex6 != expected {
			t.Errorf("expected color %s to equal %s, not %s\n", name, expected, hex6)
		}
	}
}

func TestAlternateSpelling(t *testing.T) {
	names := map[string]string{
		clrn.Aqua:           clrn.Cyan,
		clrn.DarkGray:       clrn.DarkGrey,
		clrn.DarkSlateGray:  clrn.DarkSlateGrey,
		clrn.DimGray:        clrn.DimGrey,
		clrn.Fuchsia:        clrn.Magenta,
		clrn.Gray:           clrn.Grey,
		clrn.LightGray:      clrn.LightGrey,
		clrn.LightSlateGray: clrn.LightSlateGrey,
		clrn.SlateGray:      clrn.SlateGrey,
	}

	for clrA, clrB := range names {
		a, b := clr.FromName(clrA).String(), clr.FromName(clrB).String()
		if a != b {
			t.Errorf("expected color %s (%s) to equal %s (%s)\n", clrA, a, clrB, b)
		}
	}

	if clr.Aqua.String() != clr.Cyan.String() {
		t.Error("expected clr.Aqua to equal clr.Cyan\n")
	}

	if clr.Fuchsia.ToName() != clr.Magenta.ToName() {
		t.Error("expected clr.Fuchsia to equal clr.Magenta\n")
	}
}
