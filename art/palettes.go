package art

import "image/color"

var ShasiliyRandinsky = color.Palette{
	color.RGBA{237, 229, 200, 255},
	color.RGBA{206, 198, 170, 255},
	color.RGBA{167, 200, 192, 255},
	color.RGBA{177, 39, 2, 255},
	color.RGBA{237, 189, 69, 255},
	color.RGBA{237, 189, 69, 255},
	color.RGBA{67, 130, 157, 255},
	color.RGBA{100, 132, 94, 255},
	color.RGBA{66, 157, 129, 255},
}

var paletteLookup = map[string]uint8{
	" ": 0,
	"+": 1,
	"o": 2,
	".": 3,
	"E": 4,
	"S": 5,
	"B": 6,
	"O": 7,
	"*": 8,
}
