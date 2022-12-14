package utils

import "fmt"

// 为了让其他包使用，C要大写，该函数可导出
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2

	default:
		fmt.Println("操作符号错误")
	}
	return res
}
