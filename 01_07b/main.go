package main

import (
	"flag"
	"log"
)

type operatorType int

const (
	openBracket operatorType = iota
	closedBracket
	nonBracket
)

var bracketPairs = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func getOperatorType(elem rune) operatorType {
	for leftPair, rightPair := range bracketPairs {
		switch elem {
		case leftPair:
			return openBracket
		case rightPair:
			return closedBracket
		}
	}

	return nonBracket
}

type stack struct {
	elements []rune
}

func (s *stack) Push(elem rune) {
	s.elements = append(s.elements, elem)
}

func (s *stack) Pop() *rune {
	if len(s.elements) == 0 {
		return nil
	}

	n := len(s.elements) - 1
	last := s.elements[n]
	s.elements = s.elements[:n]
	return &last
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	stk := stack{elements: []rune{}}
	for _, char := range expr {
		switch getOperatorType(char) {
		case openBracket:
			stk.Push(char)
		case closedBracket:
			last := stk.Pop()
			if last == nil || bracketPairs[*last] != char {
				return false
			}
		}
	}

	return len(stk.elements) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
