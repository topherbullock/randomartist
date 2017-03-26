package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"regexp"

	"github.com/topherbullock/randomartist/art"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, err := readUntilMatch(reader, start)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := readUntilMatch(reader, end)
	if err != nil {
		log.Fatal(err)
	}
	var height = len(contents)
	var width = len(contents[0])
	for _, line := range contents {
		fmt.Println(string(line[1 : len(line)-1]))
	}

	image.NewPaletted(getRekt(width, height), art.ShasiliyRandinsky)
}

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

func getRekt(width int, height int) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{width, height},
	}
}
