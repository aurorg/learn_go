package main

import "fmt"

func main() {
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}

	for i, v := range slice {
		fmt.Printf("i=%v v=%v \n", i, v)
	}
	
	//继续切片
	slice2 := slice[1:2]
	slice2[0] = 100

	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	//append
	var slice3 []int = []int{100, 200, 300}
	slice3 = append(slice3, 400, 500, 600)
	fmt.Println("slice3", slice3)

	slice3 = append(slice3, slice3...)
	fmt.Println("slice", slice3)

	//copy
	fmt.Println()
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	fmt.Println(slice4)
	fmt.Println(slice5)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4)
	fmt.Println("slice5=", slice5)

}
