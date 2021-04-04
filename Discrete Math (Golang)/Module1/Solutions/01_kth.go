package main

import (
	"fmt"
	"math"
)

func iPow10(n int64) int64 {
	return int64(math.Pow10(int(n)))
}

func calculateDigit(num int64) int64 {
	if num < 9 {
		return num + 1
	}

	var i, result int64
	i = 1

	for temp := num - 8; temp > 0; i++ {
		num = temp
		temp = num - 9*iPow10(i)*(i+1)
	}
	
	if num/i == 0 {
		result = iPow10(i - 1)
	} else {
		result = iPow10(i-1) + num%i + num/i - 1
	}
	
	for num = i - num%i; num > 0 && num%i != 0; num-- {
		result /= 10
	}

	return result % 10
}

func main() {
	var num int64
	fmt.Scan(&num)
	fmt.Println(calculateDigit(num))
}
