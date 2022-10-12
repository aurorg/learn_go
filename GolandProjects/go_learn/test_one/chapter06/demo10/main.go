package main

import (
	"fmt"
	"strings"
)

// 闭包可以理解为类
func AddUpper() func(int) int {
	var n int = 10 //n初始化一次，每次累加
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str=", str)
		return n
	}
}

// 返回的匿名函数和makeSuffix(suffix string)的suffix变量和返回的函数组合成一个闭包，因为返回的函数引用到suffix这个变量
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 不使用闭包，普通方法
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}
func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13闭包
	fmt.Println(f(3)) //16

	f1 := makeSuffix(".jpg") //使用闭包只用传一次后缀
	fmt.Println("文件处理后=", f1("winter"))
	fmt.Println("文件处理后=", f1("hahaha.jpg"))

	//普通方法每一次都要传入后缀
	fmt.Println("文件处理后=", makeSuffix2("jpg", "winter"))
	fmt.Println("文件处理后=", makeSuffix2("jpg", "hahaha.jpg"))
}
