package main

import "fmt"

// 全局变量
// 在go中函数外部定义变量就是全局变量
var n1 = 100
var n2 = 200
var name = "jack"

// 可以一次性声明
var (
	n3    = 300
	n4    = 900
	name2 = "mary"
)

func main() {
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	var na1, name, na3 = 100, "tom", 688
	fmt.Println("na1=", na1, "name=", name, "na3=", na3)

	ns1, name2, ns3 := 100, "tom", 678
	fmt.Println("ns1=", ns1, "name2=", name2, "ns3=", ns3)

}
