package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/chanchal1987/aoc2024go/utils"
)

func main() {
	fi := utils.Must(os.Open(utils.DataFile()))
	defer fi.Close()

	approxLen := utils.Must(fi.Stat()).Size() / 5
	left := make([]int, 0, approxLen)
	right := make([]int, 0, approxLen)
	l := 0

	for scanner := bufio.NewScanner(fi); scanner.Scan(); {
		ids := strings.Fields(scanner.Text())
		if len(ids) != 2 {
			panic("invalid line: " + scanner.Text())
		}

		left = append(left, utils.Must(strconv.Atoi(ids[0])))
		right = append(right, utils.Must(strconv.Atoi(ids[1])))
		l++
	}

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
