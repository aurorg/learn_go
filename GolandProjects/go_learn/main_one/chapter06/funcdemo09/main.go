package main

import "fmt"

var (
	Func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {

	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(20, 30)

	fmt.Println(res1)

	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res2 := a(10, 20)
	fmt.Println(res2)

	res4 := Func1(4, 9)
	fmt.Println(res4)
}
