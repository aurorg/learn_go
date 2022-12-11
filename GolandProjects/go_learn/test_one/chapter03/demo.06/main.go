package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println("i=", i)

	//测试int8的范围 -128~127
	var j int8 = -128
	fmt.Println("j=", j)

	var c byte = 'a'
	fmt.Println(c)

	//如何让查看一个变量的数据类型
	//fmt.Printf() 可以用作格式化输出
	var n1 = 100
	fmt.Printf("n1的类型是：%T \n", n1)

	//如何查看一个变量的占用的字节大小和数据类型
	var n2 int64 = 10
	fmt.Printf("n2的类型是：%T ,n2占用的字节数是%d ", n2, unsafe.Sizeof(n2))

	//保小不保大

}
