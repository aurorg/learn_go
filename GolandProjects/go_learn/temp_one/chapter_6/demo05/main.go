package main

import "fmt"

func getsum(n1 int, n2 int) int {

	return n1 + n2
}

func myfunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	a := getsum
	fmt.Printf("a的类型是%T，getsum类型是%T\n", a, getsum)

	result := a(10, 40)
	fmt.Println("result = ", result)

	res2 := myfunc(getsum, 50, 60)
	fmt.Println("res2= ", res2)

}
