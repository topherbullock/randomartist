package art

import "image/color"

var Palettes = map[string]color.Palette{
	"ShasiliyRandinsky": color.Palette{
		color.RGBA{237, 229, 200, 255},
		color.RGBA{206, 198, 170, 255},
		color.RGBA{167, 200, 192, 255},
		color.RGBA{177, 39, 2, 255},
		color.RGBA{237, 189, 69, 255},
		color.RGBA{237, 189, 69, 255},
		color.RGBA{67, 130, 157, 255},
		color.RGBA{100, 132, 94, 255},
		color.RGBA{66, 157, 129, 255},
		color.RGBA{255, 255, 255, 255},
	},
	"SharryNight": color.Palette{
		color.RGBA{29, 41, 91, 255},
		color.RGBA{219, 214, 84, 255},
		color.RGBA{22, 27, 56, 255},
		color.RGBA{60, 123, 166, 255},
		color.RGBA{214, 209, 151, 255},
		color.RGBA{50, 40, 38, 255},
		color.RGBA{121, 164, 157, 255},
		color.RGBA{56, 118, 165, 255},
		color.RGBA{239, 225, 152, 255},
		color.RGBA{57, 138, 185, 255},
	},
	"RandyWarhol": color.Palette{
		color.RGBA{244, 248, 9, 255},
		color.RGBA{166, 208, 63, 255},
		color.RGBA{61, 43, 154, 255},
		color.RGBA{249, 139, 0, 255},
		color.RGBA{245, 0, 129, 255},
		color.RGBA{247, 119, 0, 255},
		color.RGBA{247, 119, 0, 255},
		color.RGBA{37, 30, 32, 255},
		color.RGBA{248, 253, 173, 255},
		color.RGBA{106, 208, 249, 255},
	},
}

var paletteLookup = map[string]uint8{
	" ": 0,
	"|": 9,
	"+": 1,
	"o": 2,
	".": 3,
	"E": 4,
	"S": 5,
	"B": 6,
	"O": 7,
	"*": 8,
}
