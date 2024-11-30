package parser

import (
	"bufio"
	"io"
	"os"
)

func ParseByLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out, nil
}

func ParseByRunes(path string) ([]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := bufio.NewReader(file)

	var out []rune
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		} else {
			out = append(out, c)
		}
	}

	return out, nil
}
