package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b byte %T b=%v", b, b)

	var str2 string = "123456"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2, 10, 64)
	n2 = int(n1)
	fmt.Printf("n1 type %T ,n1=%v\n", n1, n1)
	fmt.Printf("n2 type %T,n2=%v\n", n2, n2)
}
