package main

import (
	"awesomeProject/main_one/chapter06/funcdemo01/utils"
	"fmt"
)

//func cal(n1 float64, n2 float64, operator byte) float64 {
//	var res float64
//	switch operator {
//	case '+':
//		res = n1 + n2
//	case '-':
//		res = n1 - n2
//	case '*':
//		res = n1 * n2
//	case '/':
//		res = n1 / n2
//
//	default:
//		fmt.Println("操作符号错误")
//	}
//	return res
//}

func main() {

	result := utils.Cal(2, 3, '+')
	//result := cal(1, 2, '-')
	fmt.Println("result=", result)
}
