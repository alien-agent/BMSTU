package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var length, index int

var variablesList = make([]string, 0)
var variablesMap = make(map[string]int)

var lexemsArray = make([]Lexem, 0)
var actions = make([]string, 0)

const (
	ERROR Tag = 1 << iota
	NUMBER
	VAR
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
)

type Tag int

type Lexem struct {
	Tag
	Image string
}

func throwError() {
	fmt.Println("error")
	os.Exit(0)
}

func isNum(l uint8) bool {
	return l >= '0' && l <= '9'
}

func isAlpha(l uint8) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z')
}

func isSpaceSymbol(l uint8) bool {
	return l == ' ' || l == '\t' || l == '\n'
}

func getNumber(expression string, index int) int {
	for index < length && isNum(expression[index]) {
		index++
	}
	return index
}

func getVariable(expr string, index int) int {
	for index < length && (isNum(expr[index]) || isAlpha(expr[index])) {
		index++
	}
	return index
}

func sliceContains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

type Stack struct {
	data []int
	top  int
}

func makeStack() Stack {
	return Stack{make([]int, 10000), 0}
}

func (s *Stack) Push(value int) {
	s.data[s.top] = value
	s.top++
}

func (s *Stack) Pop() int {
	s.top--
	return s.data[s.top]
}

func lexer(expr string, lexems chan Lexem) {
	var closeIndex int
	var el Lexem
	var operationsMap = map[uint8]Tag{'(': LPAREN, ')': RPAREN, '*': MUL, '+': PLUS, '-': MINUS, '/': DIV}

	for i := 0; i < length; i++ {
		tag, isOperationOrParen := operationsMap[expr[i]]
		if isOperationOrParen {
			el.Tag = tag
			el.Image = expr[i : i+1]
		} else if isSpaceSymbol(expr[i]) {
			continue
		} else if isNum(expr[i]) {
			closeIndex = getNumber(expr, i)
			el.Tag = NUMBER
			el.Image = expr[i:closeIndex]
			i = closeIndex - 1
		} else if isAlpha(expr[i]) {
			closeIndex = getVariable(expr, i)
			el.Tag = VAR
			el.Image = expr[i:closeIndex]
			i = closeIndex - 1
		} else {
			el.Tag = ERROR
		}
		lexems <- el
	}
}

func performActions() {
	var intermediateResult, first, second int
	var stack = makeStack()

	for _, action := range actions {
		switch action[0] {
		case '/':
			first, second = stack.Pop(), stack.Pop()
			intermediateResult = second / first
		case '*':
			first, second = stack.Pop(), stack.Pop()
			intermediateResult = first * second
		case '+':
			first, second = stack.Pop(), stack.Pop()
			intermediateResult = first + second
		case '-':
			first = stack.Pop()
			intermediateResult = -first
		default:
			if isNum(action[0]) {
				intermediateResult, _ = strconv.Atoi(action)
			} else {
				intermediateResult = variablesMap[action]
			}
		}
		stack.Push(intermediateResult)
	}

	answer := stack.Pop()
	fmt.Println(answer)
}

func parseExpression() {
	parseSandwich()
	parseButter()
}

func parseSandwich() {
	parseBread()
	parseSausage()
}

func parseSausage() {
	if index >= length {
		return
	}
	var lx = lexemsArray[index]

	if lx.Tag&(DIV|MUL) != 0 {
		index++
		parseBread()
		actions = append(actions, lx.Image)
		parseSausage()
	} else if lx.Tag&(VAR|NUMBER) != 0 {
		throwError()
	}
}

func parseButter() {
	if index >= length {
		return
	}
	var lx = lexemsArray[index]

	if lx.Tag&(PLUS|MINUS) != 0 {
		index++
		parseSandwich()
		actions = append(actions, lx.Image)
		if lx.Tag&MINUS != 0 {
			actions = append(actions, "+")
		}
		parseButter()
	} else if lx.Tag&(VAR|NUMBER) != 0 {
		throwError()
	}
}

func parseBread() {
	var lexem Lexem
	if index >= length {
		throwError()
	} else {
		lexem = lexemsArray[index]
		if lexem.Tag&(ERROR|RPAREN|PLUS|DIV|MUL) != 0 {
			throwError()
		}
	}

	if lexem.Tag&(NUMBER|VAR) != 0 {
		index++
		actions = append(actions, lexem.Image)
		if lexem.Tag&VAR != 0 {
			if !sliceContains(variablesList, lexem.Image) {
				variablesList = append(variablesList, lexem.Image)
			}
		}
	}
	if lexem.Tag&MINUS != 0 {
		index++
		parseBread()
		actions = append(actions, lexem.Image)
	}
	if lexem.Tag&LPAREN != 0 {
		index++
		parseExpression()
		if index < length {
			lexem = lexemsArray[index]
			index++
		}
		if lexem.Tag&RPAREN == 0 {
			throwError()
		}
	}
}

func main() {
	var expression string
	var temp string

	reader := bufio.NewReader(os.Stdin)
	expression, _ = reader.ReadString('\n')
	length = len(expression)

	lexems := make(chan Lexem, length)
	lexer(expression, lexems)
	close(lexems)
	for x := range lexems {
		lexemsArray = append(lexemsArray, x)
	}

	length = len(lexemsArray)
	parseExpression()
	for _, variableName := range variablesList {
		temp, _ = reader.ReadString('\n')
		variablesMap[variableName], _ = strconv.Atoi(strings.TrimSpace(temp))
	}
	performActions()
}