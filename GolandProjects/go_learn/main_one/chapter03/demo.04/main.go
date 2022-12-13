package main

import "fmt"

// 变量使用的注意事项
func main() {
	//该区域的数据值可以在同一类型范围内不断变化
	var i int = 10
	i = 30
	i = 50
	fmt.Println("i=", i)

	var a int = 1
	var b int = 2
	var c = a + b
	fmt.Println("c=", c)

	var str1 = "hello "
	var str2 = "world"
	var s = str1 + str2 //拼接
	fmt.Println(s)
}

//
//i= 50
//c= 3
//hello world
