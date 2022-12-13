package main

import (
	"awesomeProject/main_one/chapter06/funcinit08/utils"
	"fmt"
)

var age = test()

func test() int {
	fmt.Println("test()")
	return 90
}

func init() {
	fmt.Printf("init....\n")
}
func main() {
	fmt.Printf("main.....\n")

	fmt.Println("Age=", utils.Age)

}
