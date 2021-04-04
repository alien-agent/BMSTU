package main

import (
	"fmt"
	"math"
	"sort"
)

func upperBound(x int) int {
	return int(math.Pow10(int(math.Log10(float64(x))) + 1))
}

func score(a, b int) int {
	if a == 0 {
		return -283198931
	}
	if b == 0 {
		return 12831923
	}
	return a*(upperBound(b)-1) + b*(1-upperBound(a))
}

func main() {
	var numbersCount int
	fmt.Scan(&numbersCount)
	var numbers = make([]int, numbersCount)

	for i := 0; i < numbersCount; i++ {
		fmt.Scan(&numbers[i])
	}

	sort.Slice(numbers, func(i, j int) bool { return score(numbers[i], numbers[j]) > 0 })

	for i := range numbers {
		fmt.Print(numbers[i])
	}
}
