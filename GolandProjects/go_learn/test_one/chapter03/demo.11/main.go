package main

import "fmt"

func main() {
	//字符串一旦赋值就不能修改
	var address string = "北京 hello"
	fmt.Println(address)

	str2 := "abc\nabc"
	fmt.Println(str2)

	//使用反引号``
	str3 := `春色凤城来，寒梅逼岁开。条风初入树，缥雪渐侵苔。
粉逐莺衣散，香黏蝶翅回。陇头人未返，急管莫频催。`
	fmt.Println(str3)

	//字符串连接方式
	str4 := "hello" + "xupt"
	fmt.Println(str4)
}

//输出：
//北京 hello
//abc
//abc
//春色凤城来，寒梅逼岁开。条风初入树，缥雪渐侵苔。
//粉逐莺衣散，香黏蝶翅回。陇头人未返，急管莫频催。
//helloxupt
