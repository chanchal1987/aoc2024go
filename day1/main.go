package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/chanchal1987/aoc2024go/utils"
)

func main() {
	fi := utils.DataFile()
	defer fi.Close()

	approxLen := min(utils.Must(fi.Stat()).Size()/6, 1024)
	left := make([]int, 0, approxLen)
	right := make([]int, 0, approxLen)

	for scanner := bufio.NewScanner(fi); scanner.Scan(); {
		ids := strings.Fields(scanner.Text())
		if len(ids) != 2 {
			panic("invalid line: " + scanner.Text())
		}

		left = append(left, utils.Must(strconv.Atoi(ids[0])))
		right = append(right, utils.Must(strconv.Atoi(ids[1])))
	}

	l := len(left)
	left = left[:l:l]
	right = right[:l:l]
	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	similarityScore := 0

	for i := 0; i < l; i++ {
		distance += utils.AbsInt(left[i] - right[i])

		count := 0
		for pos := sort.SearchInts(right, left[i]); pos < l && right[pos] == left[i]; pos++ {
			count++
		}

		similarityScore += left[i] * count
	}

	fmt.Println("Distance:", distance)
	fmt.Println("Similarity Score:", similarityScore)
}
