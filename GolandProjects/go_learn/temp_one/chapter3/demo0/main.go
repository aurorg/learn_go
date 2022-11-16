package main

import "fmt"

func main() {

	var num1 int = 99
	//var num2 float64 =23.23445
	//var b bool =true
	//var mychar byte ='H'
	var str string
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%v\n", str, str)

}
