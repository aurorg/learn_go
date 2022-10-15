package main

import "fmt"

func main() {
	var heroes [3]string = [...]string{"ll", "ooo", "pppp"}
	//常规遍历
	//forrange遍历

	for i, v := range heroes {
		fmt.Printf("i=%v\n  v=%v\n", i, v)
	}
}
