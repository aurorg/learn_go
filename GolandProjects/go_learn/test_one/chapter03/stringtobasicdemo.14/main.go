package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool

	//b,_=strconv.ParseBool(str)
	//1.strconv.ParseBool(str)函数会返回两个值（value bool,err error）
	//2.目前只需要获取value bool，不需要获取err，所以采用_忽略

	b, _ = strconv.ParseBool(str)
	fmt.Printf("b typr %T b=%v", b, b)

	var str2 string = "123456978"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T ,n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T ,n2=%v\n", n2, n2)
}
