package main

import "fmt"

func main() {
	var price float32 = 89.12
	fmt.Println("price=", price)

	var num1 float32 = -0.00089
	var num2 float64 = -7809556.08
	fmt.Println("num1=", num1, "num2=", num2)

	var num3 float32 = -1233.3563454543
	var num4 float64 = -1233.7832532653
	fmt.Println("num3=", num3, "num4 =", num4)

	//golang的浮点型默认声明为float64
	var num5 = 1.1
	fmt.Printf("num5的数据类型是%T \n", num5)

	num6 := 5.12
	num7 := .123 //=>0.123
	fmt.Println("num6=", num6, "num7=", num7)
}
