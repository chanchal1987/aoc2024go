package matrix

import (
	"fmt"
	"reflect"
	"strings"
)

type EqF[T any] func(T, T) bool
type Matrix[T any] [][]T

func (m Matrix[T]) Clone() Matrix[T] {
	clone := make(Matrix[T], len(m))

	for i, row := range m {
		clone[i] = make([]T, len(row))
		copy(clone[i], row)
	}

	return clone
}

func (m Matrix[T]) String() string {
	var out strings.Builder

	write := func(frmt string) {
		for _, row := range m {
			for _, col := range row {
				out.WriteString(fmt.Sprintf("%"+frmt+" ", col))
			}
			out.WriteRune('\n')
		}
	}

	switch reflect.TypeFor[T]().Kind() {
	case reflect.String:
		write("q")
	case reflect.Uint8:
		frmt := "c"

	outer:
		for _, row := range m {
			for _, col := range row {
				if col := uint8(reflect.ValueOf(col).Uint()); col < 0x20 || col > 0x7E {
					frmt = "X"
					break outer
				}
			}
		}

		write(frmt)
	default:
		write("v")
	}

	return out.String()
}

func (m Matrix[T]) Get(p Position) T {
	return m[p.X][p.Y]
}

func (m Matrix[T]) Set(p Position, v T) T {
	old := m[p.X][p.Y]
	m[p.X][p.Y] = v
	return old
}

func (m Matrix[T]) Find(eq EqF[T], vs ...T) map[Position]T {
	if len(vs) == 0 {
		return nil
	}

	out := make(map[Position]T)

	for i, row := range m {
		for j, col := range row {
			for _, v := range vs {
				if eq(v, col) {
					out[Position{i, j}] = col
				}
			}
		}
	}

	return out
}

func (m Matrix[T]) Count(eq EqF[T], vs ...T) int {
	if len(vs) == 0 {
		return 0
	}

	count := 0

	for _, row := range m {
		for _, col := range row {
			for _, v := range vs {
				if eq(v, col) {
					count++
				}
			}
		}
	}

	return count
}

func (m Matrix[T]) IsEqualCols() bool {
	if len(m) < 2 {
		return true
	}

	cols := len(m[0])

	for _, row := range m[1:] {
		if len(row) != cols {
			return false
		}
	}

	return true
}

func (m Matrix[T]) Sub(buf Matrix[T], Min, Max Position) Matrix[T] {
	if !m.IsEqualCols() {
		panic("matrix must have equal number of columns for sub-matrix")
	}

	Min.SetMinX(0)
	Min.SetMinY(0)
	Max.SetMaxX(len(m))
	Max.SetMaxY(len(m[0]))

	if Min.X >= Max.X || Min.Y >= Max.Y {
		return nil
	}

	if len(buf) < Max.X-Min.X {
		buf = make(Matrix[T], Max.X-Min.X)
	}

	for i := Min.X; i < Max.X; i++ {
		buf[i-Min.X] = m[i][Min.Y:Max.Y]
	}

	return buf
}

func (m Matrix[T]) FlipHorizontally() {
	for _, row := range m {
		for i := 0; i < len(row)/2; i++ {
			row[i], row[len(row)-i-1] = row[len(row)-i-1], row[i]
		}
	}
}

func (m Matrix[T]) FlipVertically() {
	for i := 0; i < len(m)/2; i++ {
		m[i], m[len(m)-i-1] = m[len(m)-i-1], m[i]
	}
}

func (m Matrix[T]) Transpose() Matrix[T] {
	if len(m) == 0 {
		return m.Clone()
	}

	if !m.IsEqualCols() {
		panic("matrix must have equal number of columns for transpose")
	}

	out := make(Matrix[T], len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		out[i] = make([]T, len(m))
		for j := 0; j < len(m); j++ {
			out[i][j] = m[j][i]
		}
	}

	return out
}

func (m Matrix[T]) Diagonal(Min int) Matrix[T] {
	if len(m) < 2 {
		return m.Clone()
	}

	if !m.IsEqualCols() {
		panic("matrix must have equal number of columns for diagonal")
	}

	if Min < 1 {
		Min = 1
	}

	l := len(m) + len(m[0]) - (Min * 2) + 1
	if l < 1 {
		return nil
	}

	out := make(Matrix[T], l)

	for i := 0; i < l; i++ {
		for j := Min - 1 + i; j >= 0; j-- {
			if j >= len(m) || Min-1+i-j >= len(m[j]) {
				continue
			}

			out[i] = append(out[i], m[j][Min-1+i-j])
		}
	}

	return out
}
