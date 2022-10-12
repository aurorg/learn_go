package main

import "fmt"

// 指针的练习
func main() {
	//基本数据类型在内存中的布局
	var i int = 10
	//i的地址
	fmt.Println("i的地址=", &i)

	var ptr *int = &i
	//ptr是一个指针变量
	//ptr的类型是 *int
	//ptr本身的值是&i
	fmt.Println("ptr=%v", ptr)
	fmt.Printf("ptr指向的值=%v", *ptr)
}
