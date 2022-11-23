package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// operators is the map of legal operators and their functions
var operators = map[string]func(x, y float64) float64{
	"+": func(x, y float64) float64 { return x + y },
	"-": func(x, y float64) float64 { return x - y },
	"*": func(x, y float64) float64 { return x * y },
	"/": func(x, y float64) float64 { return x / y },
}

// parseOperand parses a string to a float64
func parseOperand(op string) (*float64, error) {
	parsedOp, err := strconv.ParseFloat(op, 64)
	if err != nil {
		return nil, fmt.Errorf("It was not able to parse %v", op)
	}
	return &parsedOp, nil
}

// calculate returns the result of a 2 operand mathematical expression
func calculate(expr string) (*float64, error) {
	ops := strings.Fields(expr)
	if len(ops) != 3 {
		return nil, fmt.Errorf("Invalid calculation, waited 3 operators, got: %d", len(ops))
	}

	left, err := parseOperand(ops[0])
	if err != nil {
		return nil, err
	}
	right, err := parseOperand(ops[2])
	if err != nil {
		return nil, err
	}
	f, exists := operators[ops[1]]
	if !exists {
		return nil, fmt.Errorf("Operator %s is unknown", ops[1])
	}

	result := f(*left, *right)
	return &result, nil
}

func main() {
	expr := flag.String("expr", "",
		"The expression to calculate on, separated by spaces.")
	flag.Parse()
	result, err := calculate(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s = %.2f\n", *expr, *result)
}
