package main

import "fmt"

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("n1=", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getsum sum=", sum)
	//当函数有return时，就是将结果返回给调用者
	//谁调用我，就返回给谁
	return sum

}

func main() {
	//调用test
	n1 := 10
	test(n1)
	fmt.Println("n1~=", n1)

	sum := getSum(10, 20)
	fmt.Println("main sum=", sum)

}
