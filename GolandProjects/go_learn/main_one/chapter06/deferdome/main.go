package main

import "fmt"

func sum(n1 int, n2 int) int {
	//当执行到defer时，会将后面的语句压入到独立的栈中（defer栈）
	//当函数执行完毕后，再从defer栈中，按照先入后出的方式出栈
	//!!在defer将语句放入到栈时，也会将相关的值拷贝同时入栈
	//defer主要用在，当函数执行完毕后，可以及时的释放函数创建的资源
	defer fmt.Println("ok1 n1=", n1) //n1=10
	defer fmt.Println("ok2 n2=", n2) //n2=20

	//
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}
