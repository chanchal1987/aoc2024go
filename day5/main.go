package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/chanchal1987/aoc2024go/utils"
)

type Rule struct{ Left, Right int }

func CompareRule(a, b Rule) int {
	if a.Left == b.Left {
		return a.Right - b.Right
	}

	return a.Left - b.Left
}

func SortPages(pages []int, rules []Rule) (changed bool) {
	slices.SortFunc(pages, func(i, j int) int {
		_, found := slices.BinarySearchFunc(rules, Rule{i, j}, CompareRule)
		if found {
			changed = true
			return -1
		}

		_, found = slices.BinarySearchFunc(rules, Rule{j, i}, CompareRule)
		if found {
			return 1
		}

		return 0
	})

	return
}

func main() {
	fi := utils.DataFile()
	defer fi.Close()

	var rules []Rule

	total := 0
	total2 := 0

	scanPages := false
	for scanner := bufio.NewScanner(fi); scanner.Scan(); {
		str := strings.TrimSpace(scanner.Text())
		if str == "" {
			slices.SortFunc(rules, CompareRule)
			scanPages = true
			continue
		}

		if !scanPages { // rules
			rp := strings.Split(str, "|")
			if len(rp) != 2 {
				panic("expected valid rule")
			}

			rules = append(rules, Rule{utils.Must(strconv.Atoi(rp[0])), utils.Must(strconv.Atoi(rp[1]))})
		} else {
			pp := strings.Split(str, ",")
			if len(pp) < 1 {
				panic("expected valid pages line")
			}

			mid := 0
			mid2 := 0
			changed := false

			if len(pp) != 1 {
				pages := make([]int, len(pp))

				for i, p := range pp {
					pages[i] = utils.Must(strconv.Atoi(p))
				}

				mid = pages[len(pages)/2]
				changed = SortPages(pages, rules)
				mid2 = pages[len(pages)/2]
			} else {
				mid = utils.Must(strconv.Atoi(pp[0]))
			}

			if !changed {
				total += mid
			} else {
				total2 += mid2
			}
		}
	}

	fmt.Println("Total:", total)
	fmt.Println("Total 2:", total2)
}
