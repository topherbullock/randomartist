package art

import (
	"bufio"
	"regexp"
)

func ReadRandart(reader *bufio.Reader) (*[][]byte, error) {
	_, err := readUntilMatch(reader, start)
	if err != nil {
		return nil, err
	}

	contents, err := readUntilMatch(reader, end)
	if err != nil {
		return nil, err
	}

	return &contents, nil
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
