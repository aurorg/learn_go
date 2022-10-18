package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["1"] = "111"
	a["2"] = "222"
	fmt.Println(a)
}
