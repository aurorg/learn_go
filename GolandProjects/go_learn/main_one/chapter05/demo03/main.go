package main

import "fmt"

func main() {
	//传统方式
	var str string = "hello,world!"

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}
	//输出
	//h
	//e
	//l
	//l
	//o
	//,
	//w
	//o
	//r
	//l
	//d
	//!

	str = "abc-ok"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c \n", index, val)
	}

	//输出
	//index=0, val=a
	//index=1, val=b
	//index=2, val=c
	//index=3, val=-
	//index=4, val=o
	//index=5, val=k





}
