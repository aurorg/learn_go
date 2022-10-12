package main

import "fmt"

func main() {
	//字符串一旦赋值就不能修改
	var address string = "北京 hello"
	fmt.Println(address)

	str2 := "abc\nabc"
	fmt.Println(str2)

	//使用反引号``
	str3 := `package main

import "fmt"

func main() {
	//字符串一旦赋值就不能修改
	var address string = "北京 hello"
	fmt.Println(address)

	str2 := "abc\nabc"
	fmt.Println(str2)`
	fmt.Println(str3)
}
