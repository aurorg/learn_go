package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1, 44, 22, 6, 9}
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)
	fmt.Println("slice 的元素是= ", slice)
	fmt.Println("slice的元素个数 = ", len(slice))
	fmt.Println("slice的容量 =", cap(slice))

	fmt.Println()

	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	fmt.Println(slice2)
	fmt.Println("slice2的size=", len(slice2))
	fmt.Println("slice2的cap= ", cap(slice2))

	fmt.Println()
	var strslice []string = []string{"tom", "jack", "mary"}
	fmt.Println("strslice=", strslice)
	fmt.Println("strslice size=", len(strslice))
	fmt.Println("strslice cap=", cap(strslice))

}
