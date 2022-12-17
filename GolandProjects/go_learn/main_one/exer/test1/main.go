package main

import "fmt"

func main() {

	mes := map[string]int{
		"zhangsan": 11,
		"lisi":     12,
		"wangwu":   13,
	}
	name := "zhangsan"
	v, ok := mes[name]
	if ok {
		fmt.Printf("The user mes %q: %d\n", name, v)
	} else {
		fmt.Println("fail")
	}

}
