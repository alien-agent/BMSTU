package main

import (
	"fmt"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

type Fraction struct {
	Top, Bottom int
}

func (f Fraction) String() string {
	f.Reduce()
	return fmt.Sprintf("%d/%d", f.Top, f.Bottom)
}

func (f *Fraction) Reduce() Fraction {
	gcd := GCD(f.Top, f.Bottom)
	f.Top /= gcd
	f.Bottom /= gcd
	if f.Bottom < 0 { // Приводим дроби вида -a/-b и a/-b к нормальному виду
		f.Top *= -1
		f.Bottom *= -1
	}
	return *f
}

func (f Fraction) Add(m Fraction) Fraction {
	if m.Top == 0 {
		return f
	} else if f.Top == 0 {
		return m
	}
	if f.Bottom == m.Bottom {
		f.Top += m.Top
	} else {
		f.Top *= m.Bottom
		f.Top += m.Top * f.Bottom
		f.Bottom *= m.Bottom
	}

	return f.Reduce()
}

func (f Fraction) MultiplyByNumber(m int) Fraction {
	f.Top *= m
	return f.Reduce()
}

func (f Fraction) MultiplyByFraction(m Fraction) Fraction {
	f.Top *= m.Top
	f.Bottom *= m.Bottom
	return f.Reduce()
}

type Matrix struct {
	Rows                 [][]Fraction
	RowsCount, ColsCount int
}

func (m Matrix) String() string {
	result := make([]string, m.RowsCount)
	for i := 0; i < m.RowsCount; i++ {
		result = append(result, fmt.Sprintf("%4v", m.Rows[i]))
	}
	return strings.Join(result, "\n")
}

func ReadMatrix(size int) (result Matrix) {
	result.RowsCount = size
	result.ColsCount = size + 1
	result.Rows = make([][]Fraction, result.RowsCount)
	var number int

	for i := 0; i < result.RowsCount; i++ {
		result.Rows[i] = make([]Fraction, result.ColsCount)
		for j := 0; j < result.ColsCount; j++ {
			fmt.Scanf("%d", &number)
			result.Rows[i][j] = Fraction{Top: number, Bottom: 1}
			if j == result.ColsCount-1 {
				fmt.Scanf("%d", &number)
			}
		}
	}
	return result
}

/*func (m Matrix) Sort() {
	// Сортируем строки матрицы так, чтобы на диагонали были только ненулевые элементы.
	last_incorrect := -1
	for pivotIndex := 0; pivotIndex < m.RowsCount; pivotIndex++ {
		if m.Rows[pivotIndex][pivotIndex].Top == 0 {
			minVarsCount, minVarsIndex := 100, -1
			for i := last_incorrect + 1; i < m.RowsCount; i++ {
				if m.Rows[i][pivotIndex].Top != 0 {
					varsCount := 0
					for _, val := range m.Rows[i] {
						if val.Top != 0 {
							varsCount++
						}
					}
					// Из нескольких возможных строк, имеющих ненулевое значение в нужном месте выбираем ту строку,
					// в которой больше всего нулей. Строки с большим количеством ненулевых переменных могут пригодиться
					// в другом месте.
					if varsCount < minVarsCount {
						minVarsCount = varsCount
						minVarsIndex = i
					}
				}
			}
			m.Rows[minVarsIndex], m.Rows[pivotIndex] = m.Rows[pivotIndex], m.Rows[minVarsIndex]
			last_incorrect = pivotIndex
		}
	}
}*/

func (m Matrix) IsFunny() bool { // funny == invalid
	variablesSpecified := make([]bool, m.RowsCount)
	for _, row := range m.Rows {
		zerosOnly := true
		for colIndex, val := range row {
			if colIndex == m.ColsCount-1 {
				if val.Top != 0 && zerosOnly {
					return true
				}
			} else if val.Top != 0 {
				variablesSpecified[colIndex] = true
				zerosOnly = false
			}
		}
	}
	for _, val := range variablesSpecified {
		if !val {
			return true
		}
	}
	return false
}

// Домножает строку[row] матрицы на такое число, чтобы m[row][col] стал равен 1
func (m Matrix) NormalizeElement(row, col int) {
	target := m.Rows[row][col]
	if target.Top == target.Bottom || target.Top == 0 {
		return
	}
	inverted := Fraction{target.Bottom, target.Top}
	for i := 0; i < m.ColsCount; i++ {
		m.Rows[row][i] = m.Rows[row][i].MultiplyByFraction(inverted)
	}
}

func (m Matrix) Solve() {
	for pivotIndex := 0; pivotIndex < m.RowsCount; pivotIndex++ {

		// Если опорный элемент равен 0, делаем его ненулевым, прибавляя к данной строке некоторую другую
		if m.Rows[pivotIndex][pivotIndex].Top == 0 {
			for i := pivotIndex + 1; i < m.RowsCount; i++ {
				if m.Rows[i][pivotIndex].Top != 0 {
					for colIndex := 0; colIndex < m.ColsCount; colIndex++ {
						m.Rows[pivotIndex][colIndex] = m.Rows[pivotIndex][colIndex].Add(m.Rows[i][colIndex])
					}
					break
				}
			}
			//fmt.Print(m)
		}

		m.NormalizeElement(pivotIndex, pivotIndex)
		for rowIndex := 0; rowIndex < m.RowsCount; rowIndex++ {
			if pivotIndex == rowIndex {
				continue
			}
			multiplier := m.Rows[rowIndex][pivotIndex].MultiplyByNumber(-1)
			m.Rows[rowIndex][pivotIndex].Top = 0
			for colIndex := pivotIndex + 1; colIndex < m.ColsCount; colIndex++ {
				if colIndex == pivotIndex {
					continue
				}
				m.Rows[rowIndex][colIndex] = m.Rows[rowIndex][colIndex].Add(m.Rows[pivotIndex][colIndex].MultiplyByFraction(multiplier))
			}

		}
		//fmt.Print(m)
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	matrix := ReadMatrix(n)
	if matrix.IsFunny() { // Actually funny matrices are for fun, not for solving
		fmt.Print("No solution")
	} else {
		matrix.Solve()
		if matrix.IsFunny() { // Actually funny matrices are for fun, not for solving
			fmt.Print("No solution")
		} else {
			for i := 0; i < n; i++ {
				fmt.Println(matrix.Rows[i][n])
			}
		}
	}
}
