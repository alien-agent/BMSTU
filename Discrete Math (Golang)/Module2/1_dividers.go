package main

import "fmt"

func findDivisors(n int64) []int64 {
	result := make([]int64, 0)

	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	for j := len(result) - 1; j >= 0; j-- {
		if divisor := n/result[j]; divisor * divisor != n{ // exclude perfect square duplicates
			result = append(result, divisor)
		}
	}

	return result
}

// (divisor is REALLY a divisor of number) && (number / divisor is prime)
func isDirectDivisor(number, divisor int64) bool{
	if number % divisor != 0 {
		return false
	}

	coefficient := number / divisor
	for i := int64(2); i * i <= coefficient; i++{
		if coefficient % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int64
	fmt.Scanf("%d", &n)

	divisors := findDivisors(n)

	fmt.Println("graph {")
	for i := len(divisors) - 1; i >= 0; i-- {
		fmt.Printf("\t%d\n", divisors[i])
	}

	for i := len(divisors) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if isDirectDivisor(divisors[i], divisors[j]) {
				fmt.Printf("\t%d -- %d\n", divisors[i], divisors[j])
			}
		}
	}
	fmt.Println("}")
}
