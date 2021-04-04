package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInputData() (byte, byte, []byte) {
	result := make([]byte, 0)
	reader := bufio.NewReader(os.Stdin)
	for temp, _, _ := reader.ReadRune(); temp != '\n'; temp, _, _ = reader.ReadRune() {
		result = append(result, byte(temp % 256))
	}

	a, _, _ := reader.ReadRune()
	reader.ReadRune()
	b, _, _ := reader.ReadRune()

	return byte(a % 256), byte(b % 256), result
} 

func main() {

	a, b, str := readInputData()

	minDistance := 100000
	lastAIndex, lastBIndex, diff := -100000, -100000, 0

	for currentIndex := 0; currentIndex < len(str); currentIndex++ {
		if str[currentIndex] == a {
			if lastAIndex, diff = currentIndex, currentIndex-lastBIndex; diff < minDistance {
				minDistance = diff
			}
		} else if str[currentIndex] == b {
			if lastBIndex, diff = currentIndex, currentIndex-lastAIndex; diff < minDistance {
				minDistance = diff
			}
		}
	}

	fmt.Println(minDistance - 1)
}
