package main

import "fmt"

// golang中字符类型的使用
func main() {

	var c1 byte = 'a'
	var c2 byte = '0'

	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//输出对应的字符
	fmt.Printf("c1=%c  c2=%c \n ", c1, c2)

	var c3 int = '北'
	fmt.Printf("c3=%c ", c3)
}
