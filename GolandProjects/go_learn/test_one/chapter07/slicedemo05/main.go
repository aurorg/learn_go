package main

import "fmt"

func main() {
	//string是不可变的，也就是说不能通过str[0] ='z'这样的方式来修改字符串
	//如果要修改字符串，可以先将字符串string ->[]byte 或者 []rune ->修改 -》重写转成string
	str := "helloworld"
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//转成byte之后，可以处理英文和数字，但是不能处理中文，
	//因为：【】byte 字节来处理，而一个汉字是3个字节，所以呢就会出现乱码
	//可以将string转成【】rune就好了， 【】rune按照字符来处理

	//arr1 = []rune(str)
	//arr1[0]='北'
	//str =string(arr1)
	//fmt.Println("str=",str)
}
