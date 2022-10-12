package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age) //从控制台接受一个输入

	if age > 18 {
		fmt.Println("！！！！！")
	}
}
