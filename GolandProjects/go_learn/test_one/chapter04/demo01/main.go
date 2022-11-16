package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var ispass bool
	fmt.Println("姓名")
	fmt.Scanln(&name)

	fmt.Println("年龄")
	fmt.Scanln(&age)

	fmt.Println("薪水")
	fmt.Scanln(&sal)

	//方式二
	fmt.Scanln("%s %d %f %t", &name, &age, &sal, &ispass)
	
}
