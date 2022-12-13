package main

import "fmt"

func main() {
	var key byte
	fmt.Println("清楚入一个字符：")
	fmt.Scanf("%c", &key)

	//不需要break

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")

	default:
		fmt.Println("输入有误")
	}
}
