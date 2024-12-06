package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/chanchal1987/aoc2024go/utils"
)

var MulRegExp = regexp.MustCompile(`mul\((\d+)\,(\d+)\)`)

func total(text string) int {
	sum := 0
	allMatches := MulRegExp.FindAllStringSubmatch(text, -1)

	for _, matches := range allMatches {
		if len(matches) == 0 {
			continue
		}

		if len(matches) != 3 {
			panic("invalid regexp for word \"" + text + "\"")
		}

		sum += utils.Must(strconv.Atoi(matches[1])) * utils.Must(strconv.Atoi(matches[2]))
	}

	return sum
}

func main() {
	fi := utils.DataFile()
	defer fi.Close()

	// Simple
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanWords)

	sum := 0
	for scanner.Scan() {
		sum += total(scanner.Text())
	}

	fmt.Println("Total:", sum)

	// Do & Don't
	sum = 0
	if utils.Must(fi.Seek(0, io.SeekStart)) != 0 {
		panic("unsuccessful file seek")
	}

	var dont bool
	rdr := bufio.NewReader(fi)

	text := ""
	for {
		if !dont {
			s, err := rdr.ReadString(')')
			text += s
			eof := errors.Is(err, io.EOF)
			if !strings.HasSuffix(s, "don't()") {
				if !eof {
					continue
				}
			}

			sum += total(text)
			text = ""
			dont = true
			if eof {
				break
			}
		} else {
			s, err := rdr.ReadString(')')
			if errors.Is(err, io.EOF) {
				break
			}

			if !strings.HasSuffix(s, "do()") {
				continue
			}

			dont = false
		}
	}

	fmt.Println("Conditional Total:", sum)
}
