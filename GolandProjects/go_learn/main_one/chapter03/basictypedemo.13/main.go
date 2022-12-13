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

	str = strconv.FormatInt(int64(num3), 10) //base表示几进制
	fmt.Printf("str type %T str=%q\n", str, str)

	//bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	//fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
	str = strconv.FormatFloat(num4, 'f', 10, 64) //prec控制精度（10表示10位）
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q\n", str, str)

	//strconv中的一个函数Itoa
	var num5 int = 2341
	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q\n", str, str)
}

//输出：
//str type string str=99
//str type string str="23.455600"
//str type string str="true"
//str type string str="h"
//str type string str="99"
//str type string str="23.4560000000"
//str type string str="true"
//str type string str="2341"
