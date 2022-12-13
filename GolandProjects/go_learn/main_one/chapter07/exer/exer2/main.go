package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var intArr3 [5]int
	rand.Seed(time.Now().UnixNano()) //每次生成的不一样
	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100)
	}

	fmt.Println(intArr3)

	temp := 0
	for i := 0; i < len(intArr3)/2; i++ {
		temp = intArr3[len(intArr3)-1-i]
		intArr3[len(intArr3)-1-i] = intArr3[i]
		intArr3[i] = temp
	}

	fmt.Println(intArr3)
}
