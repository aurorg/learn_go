package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//len的使用
	str := "hello"
	fmt.Println("str len =", len(str))

	str2 := "hello北京"
	//字符串遍历，如果有中文的话
	s := []rune(str2)
	for i := 0; i < len(s); i++ {
		fmt.Printf("字符 = %c\n", s[i])
	}

	//字符串转整数
	//限制输入的时候可以用
	/*
		func Atoi
		func Atoi(s string) (i int, err error)
		Atoi是ParseInt(s, 10, 0)的简写。
	*/
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转换的结果是：", n)
	}

	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Println(str)

	//字符串转byte[]
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//byte[]转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//十进制转2，8，16进制  str =strconv.FormatInt(需要转的十进制数,目标进制)
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)

	//查找子串是否在指定的字符串中
	/*
		func Contains(s, substr string) bool
		判断字符串s是否包含子串substr。
	*/
	b := strings.Contains("seafood", "foo")
	fmt.Println(b)

	//统计一个字符串中有几个指定的字串
	num := strings.Count("ceheese", "e")
	fmt.Println(num)

	//不区分大小写的字符串比较
	//==区分大小写
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b =%v\n", b)

	//返回自串在字符串第一次出现的index值，如果没有返回-1
	index := strings.Index("NLT_abcabc", "abc")
	fmt.Printf("index =%v\n", index)

	//替换
	str = strings.Replace("go go hello", "go", "go语言", 1)
	fmt.Printf("str=%v\n", str)

	//按照指定的某个字符，为分割标志，将一个字符串拆分为字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Printf("strArr=%v\n", strArr)

	//将字符串左右两边的空格去掉
	str = strings.TrimSpace(" hhhh ")
	fmt.Printf("str=%v\n", str)

	//将字符串左右两边指定的字符去掉,也可以左边或者右边
	str = strings.Trim("!hello!", "!")
	fmt.Printf("str=%v\n", str)

	//判断字符串是否以指定的字符串开头,还有结尾
	b = strings.HasPrefix("abc//lllll", "abc")
	fmt.Printf("b = %v\n", b)
}
