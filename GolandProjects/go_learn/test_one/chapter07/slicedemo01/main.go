package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1, 44, 22, 6, 9}
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是= ", slice)
	fmt.Println("slice的元素个数 = ", len(slice))
	fmt.Println("slice的容量 =", cap(slice))
	
}
