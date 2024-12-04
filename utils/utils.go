package utils

import (
	"os"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Must[T any](v T, err error) T {
	Check(err)
	return v
}

func DataFile() *os.File {
	if len(os.Args) < 2 {
		panic("no file name provided")
	}

	return Must(os.Open(os.Args[1]))
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
