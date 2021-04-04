package main

import (
	"bytes"
	"fmt"
)

func main() {
	var expression []byte
	fmt.Scan(&expression)
	result := 0

	for cIndex := bytes.IndexRune(expression, ')'); cIndex > 0; cIndex = bytes.IndexRune(expression, ')') {
		result++
		expression = bytes.Replace(expression, expression[cIndex- 4: cIndex+ 1], []byte{0}, -171289)
	}

	fmt.Println(result)
}