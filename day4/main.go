package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/chanchal1987/aoc2024go/matrix"
	"github.com/chanchal1987/aoc2024go/utils"
)

const WORD = "XMAS"

func MatchCounts(bs [][]byte) (i int) {
	for _, b := range bs {
		i += strings.Count(string(b), WORD)
	}

	return
}

func CountAll(bs [][]byte) (i int) {
	mat := matrix.Matrix[byte](bs)
	i += MatchCounts(mat)

	mat2 := mat.Transpose()
	i += MatchCounts(mat2)

	mat3 := mat.Diagonal(len(WORD))
	i += MatchCounts(mat3)

	mat.FlipHorizontally()
	i += MatchCounts(mat)

	mat2.FlipHorizontally()
	i += MatchCounts(mat2)

	mat4 := mat2.Diagonal(len(WORD))
	i += MatchCounts(mat4)

	mat3.FlipHorizontally()
	i += MatchCounts(mat3)

	mat4.FlipHorizontally()
	i += MatchCounts(mat4)

	return
}

func main() {
	fi := utils.DataFile()
	defer fi.Close()

	data := bytes.Split(utils.Must(io.ReadAll(fi)), []byte{'\n'})
	fmt.Println("XMAS Count:", CountAll(data))

	// Part 2
	var buf matrix.Matrix[byte]

	count := 0
	for i := 0; i < len(data[0])-2; i++ {
		for j := 0; j < len(data)-2; j++ {
			buf = matrix.Matrix[byte](data).Sub(buf, matrix.Position{X: j, Y: i}, matrix.Position{X: j + 3, Y: i + 3})
			if buf[1][1] == 'A' {
				if (buf[0][0] == 'M' && buf[2][2] == 'S') || (buf[0][0] == 'S' && buf[2][2] == 'M') {
					if (buf[0][2] == 'M' && buf[2][0] == 'S') || (buf[0][2] == 'S' && buf[2][0] == 'M') {
						count++
					}
				}
			}
		}
	}

	fmt.Println("X-MAS Count:", count)
}
