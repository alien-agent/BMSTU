package main

import "fmt"

var array []int

func partition(less func(i, j int) bool, swap func(i, j int), low, high int) int {
	i, j := low, low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, high)
	return i
}

func quicksort(less func(i, j int) bool, swap func(i, j int), low, high int) {
	if high <= low {
		return
	}

	j := partition(less, swap, low, high)

	quicksort(less, swap, low, j-1)
	quicksort(less, swap, j+1, high)
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	quicksort(less, swap, 0, n-1)
}

func less(i, j int) bool {
	return array[i] < array[j]
}

func swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	array = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}

	qsort(n, less, swap)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}
}
