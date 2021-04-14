package main // Hate lexical analysis :(

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	graph       = *new([]Node)
	resultQueue = Deque{}
)

type Tag int

type Token struct {
	Image string
	Tag
}

const (
	START Tag = 1 << iota
	IDENT
	CONST
	ADD
	SUB
	MUL
	DIV
	LPAREN
	RPAREN
	COMMA
	EQUALS
)

// ***************************************
// ********** UTILITY FUNCTIONS **********
// ***************************************

func reportCycle() {
	fmt.Println("cycle")
	os.Exit(0)
}

func reportSyntaxError() {
	fmt.Println("syntax error")
	os.Exit(0)
}

func isDigit(l uint8) bool {
	return l >= '0' && l <= '9'
}

func isLetter(l uint8) bool {
	return (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z')
}

func tokenize(s string) []Token {
	tokens := make([]Token, 0)
	opsMap := map[byte]Tag{'+': ADD, '-': SUB, '/': DIV, '*': MUL, '(': LPAREN, ')': RPAREN, ',': COMMA, '=': EQUALS}
	lastTag := START

	for i := 0; i < len(s); i++ {
		var image string

		if s[i] == ' ' || (s[i] == '-' && lastTag&(IDENT|CONST|RPAREN) == 0) {
			continue
		} else if tag, isOperation := opsMap[s[i]]; isOperation {
			lastTag = tag
			image = string(s[i])
		} else if isDigit(s[i]) || isLetter(s[i]) {
			var imageBuilder strings.Builder
			if isDigit(s[i]) {
				for ; i < len(s) && isDigit(s[i]); i++ {
					imageBuilder.WriteByte(s[i])
				}
				lastTag = CONST
			} else {
				for ; i < len(s) && (isDigit(s[i]) || isLetter(s[i])); i++ {
					imageBuilder.WriteByte(s[i])
				}
				lastTag = IDENT
			}
			i--
			image = imageBuilder.String()
		} else {
			reportSyntaxError()
		}
		tokens = append(tokens, Token{image, lastTag})
	}
	return tokens
}

// *************************************
// ********** STACK & DEqueue **********
// *************************************

type Stack []Token

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Last() Token {
	return (*s)[len(*s)-1]
}

func (s *Stack) Push(t Token) {
	*s = append(*s, t)
}

func (s *Stack) Pop() Token {
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

type Deque struct {
	Items []interface{}
}

func (s *Deque) Push(item interface{}) {
	temp := []interface{}{item}
	s.Items = append(temp, s.Items...)
}

func (s *Deque) Inject(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Deque) Pop() interface{} {
	defer func() {
		s.Items = s.Items[1:]
	}()
	return s.Items[0]
}

func (s *Deque) Eject() interface{} {
	i := len(s.Items) - 1
	defer func() {
		s.Items = append(s.Items[:i], s.Items[i+1:]...)
	}()
	return s.Items[i]
}

func (s *Deque) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

// ************************************
// ********** MAIN PART **********
// ************************************

type Node struct {
	Color                int
	Formula              string
	Tokens               []Token
	Idents, Dependencies []string
	To, From             map[int]bool
}

func CreateNode(formula string) Node {
	n := Node{
		Color:        0,
		Formula:      formula,
		Tokens:       tokenize(formula),
		Idents:       *new([]string),
		Dependencies: *new([]string),
		To:           make(map[int]bool),
		From:         make(map[int]bool),
	}
	n.parse()
	return n
}

func (n Node) String() string {
	return n.Formula
}

func (n *Node) parse() {
	var (
		leftPart              = make([]Token, 0)
		rightPart             = make([][]Token, 0)
		leftCount, rightCount = 0, 0
		isParsingRight        = false
	)
	rightPart = append(rightPart, make([]Token, 0, 100))

	for _, token := range n.Tokens {
		if token.Tag&EQUALS != 0 {
			isParsingRight = true
			continue
		}
		if !isParsingRight && (token.Tag&COMMA != 0) {
			leftCount++
		}
		if isParsingRight && (token.Tag&COMMA != 0) {
			rightCount++
			rightPart = append(rightPart, *new([]Token))
		}
		if isParsingRight && (token.Tag&COMMA == 0) {
			rightPart[rightCount] = append(rightPart[rightCount], token)
		}
		if !isParsingRight {
			leftPart = append(leftPart, token)
		}
	}

	for _, token := range leftPart {
		if token.Tag&(IDENT|COMMA) == 0 { // if not var or comma in left part
			reportSyntaxError()
		}
	}
	if len(rightPart) == 0 || leftCount != rightCount {
		reportSyntaxError()
	} else {
		identsCount := 0
		for _, t := range leftPart {
			if t.Tag&IDENT != 0 {
				identsCount++
			}
		}
		if identsCount != rightCount+1 {
			reportSyntaxError()
		}
	}

	n.parseLeft(leftPart)
	for _, expr := range rightPart {
		n.parseExpression(expr)
	}
}

func (n *Node) parseLeft(tokens []Token) {
	for _, t := range tokens {
		if t.Tag&IDENT != 0 {
			n.Idents = append(n.Idents, t.Image)
		}
	}
	if len(n.Idents) == 0 {
		reportSyntaxError()
	}
}

func (n *Node) parseExpression(tokens []Token) {
	if len(tokens) == 0 {
		reportSyntaxError()
	}

	q := Deque{}
	for _, v := range tokens {
		q.Inject(v)
	}
	vars := Deque{}
	ops := Stack{}
	opCnt, varCnt := 0, 0

	for !q.IsEmpty() {
		token := q.Pop().(Token)
		if token.Tag&(IDENT|CONST) != 0 {
			if token.Tag&IDENT != 0 {
				n.Dependencies = append(n.Dependencies, token.Image)
			}
			vars.Push(token)
			varCnt++
		} else if token.Tag&LPAREN != 0 {
			ops.Push(token)
		} else if token.Tag&RPAREN != 0 {
			for !ops.IsEmpty() && (ops.Last().Tag&LPAREN == 0) {
				vars.Push(ops.Pop())
			}
			if !ops.IsEmpty() && (ops.Last().Tag&LPAREN != 0) {
				ops.Pop()
				if !ops.IsEmpty() && (ops.Last().Tag&(DIV|MUL|ADD|SUB) != 0) {
					vars.Push(ops.Pop())
				}
			} else if !ops.IsEmpty() {
				reportSyntaxError()
			}
		} else {
			ops.Push(token)
			opCnt++
		}
	}

	for !ops.IsEmpty() {
		token := ops.Pop()
		if token.Tag&LPAREN != 0 {
			reportSyntaxError()
		} else {
			vars.Push(token)
		}
	}

	if varCnt-opCnt != 1 {
		reportSyntaxError()
	}
}

func depthFirstSearch(vertexIndex int) {
	graph[vertexIndex].Color = 1
	for otherIndex := range graph[vertexIndex].To {
		if graph[otherIndex].Color == 0 {
			depthFirstSearch(otherIndex)
		} else if graph[otherIndex].Color == 1 {
			reportCycle()
		}
	}

	graph[vertexIndex].Color = 2
	resultQueue.Inject(vertexIndex)
}

func main() {
	variableNodeIndex := make(map[string]int) // Index of node in graph where specific variable is specified
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; scanner.Scan(); i++ {
		newNode := CreateNode(scanner.Text())
		for _, ident := range newNode.Idents {
			if _, present := variableNodeIndex[ident]; present {
				reportSyntaxError()
			}
			variableNodeIndex[ident] = i
		}
		graph = append(graph, newNode)
	}

	for _, vertex := range graph {
		for _, dependencyKey := range vertex.Dependencies {
			if otherIndex, isDependencyPresent := variableNodeIndex[dependencyKey]; !isDependencyPresent {
				reportSyntaxError()
			} else {
				vertex.From[otherIndex] = true
			}
		}
	}
	for vertexIndex, vertex := range graph {
		for otherIndex := range vertex.From {
			graph[otherIndex].To[vertexIndex] = true
		}
	}

	for vertexIndex, vertex := range graph {
		if vertex.Color == 0 {
			depthFirstSearch(vertexIndex)
		}
	}

	for !resultQueue.IsEmpty() {
		vertexIndex := resultQueue.Eject().(int)
		fmt.Println(graph[vertexIndex])
	}
}
