package main

import "fmt"

func main() {
	var num int = 9
	fmt.Printf("num的地址是=%v\n", &num)

	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Printf("num=%v", num)
}
