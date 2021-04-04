package main

import (
	"fmt"
	"math/big"
)

// Decomposes number n into the sum of powers of 2
// Example: 13 = 1 + 4 + 8
func decompose(n int) []int {
	if n <= 2 {
		return []int{n}
	}

	powers := make([]int, 0)
	i := 1
	for i <= n {
		if i&n > 0 {
			powers = append(powers, i)
		}
		i <<= 1
	}
	return powers
}

type Matrix struct {
	rows [][]big.Int
	size int
}

func makeMatrix(array [][]int) Matrix {
	size := len(array)
	result := Matrix{make([][]big.Int, size), size}

	for i, row := range array {
		result.rows[i] = make([]big.Int, size)
		for j, val := range row {
			result.rows[i][j] = *big.NewInt(int64(val))
		}
	}

	return result
}

func (m *Matrix) MultiplyBy(other Matrix) Matrix {
	matrixSize := m.size
	var result = makeMatrix([][]int{{0, 0}, {0, 0}})

	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			for r := 0; r < matrixSize; r++ {
				result.rows[i][j].Add(&result.rows[i][j], big.NewInt(0).Mul(&m.rows[i][r], &other.rows[r][j]))
			}
		}
	}

	return result
}

func (m *Matrix) Square() Matrix {
	return m.MultiplyBy(*m)
}

func matrixMagic(power int) Matrix {
	powers := decompose(power - 1)
	matrix := makeMatrix([][]int{{1, 1}, {1, 0}})
	result := makeMatrix([][]int{{1, 0}, {0, 1}})

	for i, j := 1, 0; i <= powers[len(powers)-1]; i *= 2 {
		if i == powers[j] {
			result = result.MultiplyBy(matrix)
			j++
			if j >= len(powers) {
				break
			}
		}
		matrix = matrix.Square()
	}

	return result
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	// Should be just: for i := 0; i < n; i++ { a, b = b, a + b}
	matrix := matrixMagic(n)
	fmt.Println(matrix.rows[0][0].String())
}
