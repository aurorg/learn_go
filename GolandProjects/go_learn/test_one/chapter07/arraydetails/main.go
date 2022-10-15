package main

import "fmt"

func test01(arr [3]int) {
	arr[0] = 88
}

func test02(arr *[3]int) {
	(*arr)[0] = 88
}

func main() {
	//arr := [3]int{11, 22, 33}
	//test01(arr)
	//fmt.Println(arr)
	arr := [3]int{11, 22, 33}
	test02(&arr)
	fmt.Println(arr)
}
