package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func (person Person) test2() {
	person.Name = "jack"
	fmt.Println("test2() name=", person.Name)
}

func (p Person) speak() {
	fmt.Println(p.Name, "very good")
}

func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果=", res)
}

func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}
func main() {
	var p Person
	p.Name = "tom"
	p.test()

	p.speak()
	p.jisuan()
	p.jisuan2(20)
	res := p.getSum(10, 20)
	fmt.Println("res=", res)
}
