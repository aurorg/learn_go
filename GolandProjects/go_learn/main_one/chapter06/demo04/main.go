package main

import "fmt"

func test() {
	//var n1 int = 10
}
func test02(n1 int) {
	n1 = n1 + 10
	fmt.Println("test02 n1= ", n1)
}

func test03(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test03 n1= ", *n1)
}
func main() {

	//num := 20
	//test02(num)
	//fmt.Println("main num=", num)

	num := 20
	test03(&num)
	fmt.Println("main num=", num)
}
