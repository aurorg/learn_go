package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "hello"
	fmt.Println("str len=", len((str)))

	str2 := "hello北京"
	s := []rune(str2)
	for i := 0; i < len(s); i++ {
		fmt.Printf("字符 = %c\n", s[i])
	}

	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是：", n)
	}

	str = strconv.Itoa(12345)
	fmt.Println(str)

	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

}
