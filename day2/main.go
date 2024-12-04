package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/chanchal1987/aoc2024go/utils"
)

func main() {
	fi := utils.DataFile()
	defer fi.Close()

	safe := 0
	dSafe := 0
	for scanner := bufio.NewScanner(fi); scanner.Scan(); {
		data := strings.Fields(scanner.Text())
		if len(data) == 0 {
			panic("invalid line")
		}

		unsafe := 0
		prev := utils.Must(strconv.Atoi(data[0]))
		desc := true
		orderSet := false

		for _, datam := range data[1:] {
			datam := utils.Must(strconv.Atoi(datam))

			if !orderSet {
				if datam > prev {
					desc = false
				}

				orderSet = true
			}

			diff := datam - prev
			if desc {
				diff = -diff
			}

			if diff != 1 && diff != 2 && diff != 3 {
				unsafe++
				if unsafe > 1 {
					break
				}
			}

			prev = datam
		}

		if unsafe == 0 {
			safe++
		}

		if unsafe <= 1 {
			dSafe++
		}
	}

	fmt.Println("Safe:", safe)
	fmt.Println("Dampener Safe:", dSafe)
}
