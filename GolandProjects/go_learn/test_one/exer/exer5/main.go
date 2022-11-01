package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

func calculate(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

func main() {
	add := func(x, y int) int {
		return x + y
	}
	sub := func(x, y int) int {
		return x - y
	}

	x, y := 10, 12
	resultAdd, _ := calculate(x, y, add)
	fmt.Printf("calculate add(%v, %v) --> %v\n", x, y, resultAdd)

	a, b := 12, 10
	resultSub, _ := calculate(a, b, sub)
	fmt.Printf("calculate sub(%v, %v) --> %v\n", a, b, resultSub)
}

// go run result:
// calculate add(10, 12) --> 22
// calculate sub(12, 10) --> 2
