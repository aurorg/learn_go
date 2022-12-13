package main

import "fmt"

func main() {

	var str string = "hello,world!北京"
	str2 := []rune(str) //就是把str转成[]rune
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}

	//h
	//e
	//l
	//l
	//o
	//,
	//w
	//o
	//r
	//l
	//d
	//!
	//北
	//京

	str = "abc-ok上海"
	for index, val := range str {
		fmt.Printf("index=%d , val=%c \n", index, val)
	}

	//index=0 , val=a
	//index=1 , val=b
	//index=2 , val=c
	//index=3 , val=-
	//index=4 , val=o
	//index=5 , val=k
	//index=6 , val=上
	//index=9 , val=海

}
