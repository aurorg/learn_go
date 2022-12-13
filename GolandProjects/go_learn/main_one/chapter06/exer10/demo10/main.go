package main

import (
	"fmt"
	"strings"
)

// 闭包
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str=", str)
		return n
	}
}

// 使用闭包，返回的匿名函数和makeSuffix(suffix string)的suffix变量和返回的函数组成一个包，因为返回的函数引用到suffix这个变量了
func makeSuffix(suffix string) func(string) string {
	return func(name string) string { //匿名函数
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

	f1 := makeSuffix(".jpg")
	fmt.Println("文件处理后=", f1("winter"))
	fmt.Println("文件处理后", f1("hahaha.jpg"))

	//普通的方法每一次都要传入后缀

	fmt.Println("文件处理后=", makeSuffix2(".jpg", "winter"))
	fmt.Println("文件处理后", makeSuffix2(".jpg", "hahaha,jpg"))

}
