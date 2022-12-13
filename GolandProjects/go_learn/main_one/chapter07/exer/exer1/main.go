package main

import "fmt"

func main() {
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])
	}

	fmt.Println()

	var intArr [5]int = [...]int{1, 4, 2, 88, 99}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 0; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxval=%v maxValIndex=%v", maxVal, maxValIndex)
}
