package main

import "fmt"

// go支持可变参数
// 如果一个函数的形参列表有可变参数，则可变参数需要放在形参列表的最后
func sum(n1 int, vars ...int) int {
	sum := n1
	for i := 0; i < len(vars); i++ {
		sum += vars[i]
	}
	return sum
}
func main() {

	res4 := sum(10, 0, -1, 90, 10, 100)
	fmt.Println(res4)
}
