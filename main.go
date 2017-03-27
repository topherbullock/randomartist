package main

import (
	"bufio"
	"image"
	"log"
	"os"
	"regexp"

	"image/color"
	"image/draw"

	"image/jpeg"

	"github.com/topherbullock/randomartist/art"
)

func main() {

	// TODO : flags for input from a file
	reader := bufio.NewReader(os.Stdin)

	_, err := readUntilMatch(reader, start)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := readUntilMatch(reader, end)
	if err != nil {
		log.Fatal(err)
	}

	// TODO : flag for pallete selection
	// TODO : flag for scale
	img := paintImage(contents, art.SharryNight, 25)

	// TODO : flag for image output location
	out, err := os.Create("./out.jpg")
	if err != nil {
		log.Fatal(err)
	}

	// TODO : flag for quality
	err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})

	if err != nil {
		log.Fatal(err)
	}
}

func paintImage(contents [][]byte, palette color.Palette, scale int) *image.Paletted {
	height := len(contents) * scale
	width := len(contents[0]) * scale
	img := image.NewPaletted(getRekt(image.Point{0, 0}, width, height), palette)

	img.Pix = make([]uint8, width*height)

	for y, line := range contents {
		for x, symbol := range line {
			colorIndex := art.PaletteLookup[string(symbol)]
			box := getRekt(image.Point{x * scale, y * scale}, scale, scale)
			draw.Draw(img, box, &image.Uniform{palette[colorIndex]}, image.ZP, draw.Src)
		}
	}

	return img
}

// TODO prolly want to fix these to handle more randart outputs than the one example I've tested with
var start = regexp.MustCompile(`\+\-*\[RSA\ (\d)*\]\-*\+`)
var end = regexp.MustCompile(`\+\-*\[SHA256(\d)*\]\-*\+`)

func readUntilMatch(reader *bufio.Reader, exp *regexp.Regexp) ([][]byte, error) {
	var lines [][]byte
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			return lines, err
		}

		if !isPrefix && exp.Match(line) {
			return lines, nil
		}

		lines = append(lines, line)
	}
}

func getRekt(point image.Point, width int, height int) image.Rectangle {
	return image.Rectangle{
		Min: point,
		Max: image.Point{point.X + width, point.Y + height},
	}
}
