package main

import (
	"bufio"
	"errors"
	"image"
	"io"
	"log"
	"os"

	"image/jpeg"

	"github.com/jessevdk/go-flags"
	"github.com/topherbullock/randomartist/art"
)

type RandomArtistCmd struct {
	InputFile   string `long:"input" short:"i" description:"Source file containing randart"`
	OutputFile  string `long:"output" short:"o" default:"randart.jpg" description:"output file destination"`
	Scale       int    `long:"scale"  short:"s" default:"25" description:"scaling factor for the image file"`
	JpegQuality int    `long:"quality"  short:"q" default:"85" description:"jpeg quality for the output"`
	Palette     string `long:"palette" short:"p" default:"SharryNight" choice:"SharryNight" choice:"RandyWarhol" choice:"ShasiliyRandinsky" description:"color palette to use"`
}

func (r *RandomArtistCmd) readfile() (*[][]byte, error) {
	var (
		source io.Reader
		err    error
	)

	if r.InputFile != "" {
		source, err = os.Open(r.InputFile)
		if err != nil {
			return nil, err
		}
	} else {
		source = os.Stdin
	}

	return art.ReadRandart(bufio.NewReader(source))
}

func (r *RandomArtistCmd) paintImage(contents [][]byte) (image.Image, error) {
	palette, ok := art.Palettes[r.Palette]
	if !ok {
		return nil, errors.New("Palette not found")
	}

	return art.PaintImage(contents, palette, r.Scale), nil
}

func (r *RandomArtistCmd) saveFile(img image.Image) error {
	out, err := os.Create(r.OutputFile)
	if err != nil {
		log.Fatal(err)
	}

	return jpeg.Encode(out, img, &jpeg.Options{Quality: r.JpegQuality})
}

func main() {
	cmd := &RandomArtistCmd{}
	parser := flags.NewParser(cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	_, err := parser.Parse()

	if err != nil {
		os.Exit(1)
	}

	contents, err := cmd.readfile()
	if err != nil {
		log.Fatal(err)
	}

	img, err := cmd.paintImage(*contents)
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.saveFile(img)
	if err != nil {
		log.Fatal(err)
	}
}
