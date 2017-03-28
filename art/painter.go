package art

import (
	"image"
	"image/color"
	"image/draw"
)

func PaintImage(contents [][]byte, palette color.Palette, scale int) *image.Paletted {
	height := len(contents) * scale
	width := len(contents[0]) * scale
	img := image.NewPaletted(getRekt(image.Point{0, 0}, width, height), palette)

	img.Pix = make([]uint8, width*height)

	for y, line := range contents {
		for x, symbol := range line {
			colorIndex := paletteLookup[string(symbol)]
			box := getRekt(image.Point{x * scale, y * scale}, scale, scale)
			draw.Draw(img, box, &image.Uniform{palette[colorIndex]}, image.ZP, draw.Src)
		}
	}

	return img
}

func getRekt(point image.Point, width int, height int) image.Rectangle {
	return image.Rectangle{
		Min: point,
		Max: image.Point{point.X + width, point.Y + height},
	}
}
