package main

import "fmt"

// 结构体的声明
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明结构体的变量
	var (
		p  person
		p2 person
	)
	//为结构体变量赋值
	p.name = "卡卡"
	p.age = 24
	p.gender = "男"
	p.hobby = []string{"篮球", "羽毛球"}
	fmt.Printf("type:%T,value:%+v\n", p, p) //类型就是type定义的类型，使用%+v字段名可以打印出所有信息
	fmt.Println(p.name)
	p2.name = "zq"
	p.age = 23
	fmt.Printf("type:%T,value:%+v\n", p2, p2) //没有初始化的值声明后就是空值
	//匿名结构体，直接用var 变量名 struct{}声明一个结构体变量，多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "Hello"
	s.y = 999
	fmt.Printf("type:%T,value:%+v\n", s, s) //匿名变量的类型为struc{变量 变量类型...}
}
