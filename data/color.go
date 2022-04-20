package data

import (
	"fmt"
	"image/color"
)

type Color struct {
	Id        int
	Name      string
	RGB       []int
	Available bool
}

func StringToColor(str string) (Color, error) {
	for _, col := range Colors() {
		if col.Name == str {
			return col, nil
		}
	}
	return Color{}, fmt.Errorf("unknown color: %s", str)
}

func Colors() []Color {
	return []Color{
		White,
		LightGray,
		Gray,
		DarkGray,
		Black,
		Beige,
		Brown,
		DarkBrown,
		LightPink,
		Pink,
		Magenta,
		PalePurple,
		Purple,
		DarkPurple,
		Lavender,
		Periwinkle,
		Indigo,
		LightBlue,
		Blue,
		DarkBlue,
		LightTeal,
		Teal,
		DarkTeal,
		LightGreen,
		Green,
		DarkGreen,
		PaleYellow,
		Yellow,
		Orange,
		Red,
		DarkRed,
		Burgundy,
	}
}

var Transparent = Color{
	Id:        -1,
	Name:      "Transparent",
	RGB:       nil,
	Available: false,
}

var White = Color{
	Id:        0,
	Name:      "White",
	RGB:       []int{255, 255, 255},
	Available: true,
}

var LightGray = Color{
	Id:        1,
	Name:      "Light Gray",
	RGB:       []int{212, 215, 217},
	Available: true,
}

var Gray = Color{
	Id:        2,
	Name:      "Gray",
	RGB:       []int{137, 141, 144},
	Available: true,
}

var DarkGray = Color{
	Id:        3,
	Name:      "Dark Gray",
	RGB:       []int{81, 82, 82},
	Available: true,
}

var Black = Color{
	Id:        4,
	Name:      "Black",
	RGB:       []int{0, 0, 0},
	Available: true,
}

var Beige = Color{
	Id:        5,
	Name:      "Beige",
	RGB:       []int{255, 180, 112},
	Available: true,
}

var Brown = Color{
	Id:        6,
	Name:      "Brown",
	RGB:       []int{156, 105, 38},
	Available: true,
}

var DarkBrown = Color{
	Id:        7,
	Name:      "Dark Brown",
	RGB:       []int{109, 72, 47},
	Available: true,
}

var LightPink = Color{
	Id:        8,
	Name:      "Light Pink",
	RGB:       []int{255, 153, 170},
	Available: true,
}

var Pink = Color{
	Id:        9,
	Name:      "Pink",
	RGB:       []int{255, 56, 129},
	Available: true,
}

var Magenta = Color{
	Id:        10,
	Name:      "Magenta",
	RGB:       []int{222, 16, 127},
	Available: true,
}

var PalePurple = Color{
	Id:        11,
	Name:      "Pale Purple",
	RGB:       []int{228, 171, 255},
	Available: true,
}

var Purple = Color{
	Id:        12,
	Name:      "Purple",
	RGB:       []int{180, 74, 192},
	Available: true,
}

var DarkPurple = Color{
	Id:        13,
	Name:      "Dark Purple",
	RGB:       []int{129, 30, 159},
	Available: true,
}

var Lavender = Color{
	Id:        14,
	Name:      "Lavender",
	RGB:       []int{148, 179, 255},
	Available: true,
}

var Periwinkle = Color{
	Id:        15,
	Name:      "Periwinkle",
	RGB:       []int{106, 92, 255},
	Available: true,
}

var Indigo = Color{
	Id:        16,
	Name:      "Indigo",
	RGB:       []int{73, 58, 193},
	Available: true,
}

var LightBlue = Color{
	Id:        17,
	Name:      "Light Blue",
	RGB:       []int{73, 58, 193},
	Available: true,
}

var Blue = Color{
	Id:        18,
	Name:      "Blue",
	RGB:       []int{54, 144, 234},
	Available: true,
}

var DarkBlue = Color{
	Id:        19,
	Name:      "Dark Blue",
	RGB:       []int{36, 80, 164},
	Available: true,
}

var LightTeal = Color{
	Id:        20,
	Name:      "Light Teal",
	RGB:       []int{0, 204, 192},
	Available: true,
}

var Teal = Color{
	Id:        21,
	Name:      "Teal",
	RGB:       []int{0, 158, 170},
	Available: true,
}

var DarkTeal = Color{
	Id:        22,
	Name:      "Dark Teal",
	RGB:       []int{0, 117, 111},
	Available: true,
}

var LightGreen = Color{
	Id:        23,
	Name:      "Light Green",
	RGB:       []int{126, 237, 86},
	Available: true,
}

var Green = Color{
	Id:        24,
	Name:      "Green",
	RGB:       []int{0, 204, 120},
	Available: true,
}

var DarkGreen = Color{
	Id:        25,
	Name:      "Dark Green",
	RGB:       []int{0, 163, 104},
	Available: true,
}

var PaleYellow = Color{
	Id:        26,
	Name:      "Pale Yellow",
	RGB:       []int{255, 248, 184},
	Available: true,
}

var Yellow = Color{
	Id:        27,
	Name:      "Yellow",
	RGB:       []int{255, 214, 53},
	Available: true,
}

var Orange = Color{
	Id:        28,
	Name:      "Orange",
	RGB:       []int{255, 168, 0},
	Available: true,
}

var Red = Color{
	Id:        29,
	Name:      "Red",
	RGB:       []int{255, 69, 0},
	Available: true,
}

var DarkRed = Color{
	Id:        30,
	Name:      "Dark Red",
	RGB:       []int{190, 0, 57},
	Available: true,
}

var Burgundy = Color{
	Id:        31,
	Name:      "Burgundy",
	RGB:       []int{109, 0, 26},
	Available: true,
}

func RgbaToRgb(r uint32, g uint32, b uint32, a uint32) (int, int, int, int) {
	if a < 127 {
		return int(r / 257), int(g / 257), int(b / 257), 0
	}
	return int(r / 257), int(g / 257), int(b / 257), 255
}

func RgbToRgba(rgb []int) color.RGBA {
	return color.RGBA{
		R: uint8(rgb[0]),
		G: uint8(rgb[1]),
		B: uint8(rgb[2]),
		A: 255,
	}
}
