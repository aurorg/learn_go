package main

import (
	"fmt"
	"strconv"
)

func main() {
	//len的使用
	str := "hello"
	fmt.Println("str len =", len(str))

	str2 := "hello北京"
	//字符串遍历，如果有中文的话
	s := []rune(str2)
	for i := 0; i < len(s); i++ {
		fmt.Printf("字符 = %c\n", s[i])
	}

	//字符串转整数
	//限制输入的时候可以用
	/*
		func Atoi
		func Atoi(s string) (i int, err error)
		Atoi是ParseInt(s, 10, 0)的简写。
	*/
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是：", n)
	}

	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Println(str)

	//字符串转byte[]
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//byte[]转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

}
