package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型转成string
func main() {
	var num1 int = 99
	var num2 float64 = 23.4556
	var b bool = true
	var mychar byte = 'h'
	var str string

	//使用fmt.Sprintf方法

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n", str, str)

	//第二种方式 strconv函数
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv中的一个函数Itoa
	var num5 int = 2341
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)
}
