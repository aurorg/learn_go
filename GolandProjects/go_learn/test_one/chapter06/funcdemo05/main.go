package main

import "fmt"

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 给函数类型去一个别名
type myFunType func(int, int) int //myFun就是func(int, int) int类型的

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
func main() {
	a := getSum

	fmt.Printf("a的类型是%T，getsum类型是%T\n", a, getSum)

	result := a(10, 40) //等价于res:=getsum(10,40)
	fmt.Println("result = ", result)

	res2 := myFun2(getSum, 50, 60)
	fmt.Println("res2=", res2)

	type myInt int //给int取一个别名
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)

	fmt.Println(num1)
	fmt.Println(num2)

	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res2=", res3)

	a1, b1 := getSumAndSub(1, 2)
	fmt.Printf("a=%v b=%v\n", a1, b1)
}
