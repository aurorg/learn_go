package main

import "fmt"

func main() {
	var i int32 = 100
	//希望将i => float
	var n1 float32 = float32(i)

	fmt.Printf("i=%v\n n1=%v", i, n1)
}

//输出：
//i=100
//n1=100
