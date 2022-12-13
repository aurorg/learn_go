package main

import "fmt"

func main() {
	var key byte
	fmt.Println("清楚入一个字符：")
	fmt.Scanf("%c", &key)

	//不需要break

	switch key {
	case 'a':
		fmt.Println("周一")
	case 'b':
		fmt.Println("周二")
	case 'c':
		fmt.Println("周三")

	default:
		fmt.Println("输入有误")
	}

	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5: //case后面可以跟多个表达式，满足其一就可以了
		fmt.Println("ok1")
	default:
		fmt.Println("没有匹配到")
	}

	//switch后也可以不带表达式，类似if--else分支来使用
	var age int = 10
	switch {
	case age == 10:
		fmt.Println("10")
	case age == 20:
		fmt.Println("20")
	default:
		fmt.Println("没有匹配到")
	}

	//switch 的穿透fallthrough

	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //只能穿透一层，如果num=10，会输出ok1，因为穿透的原因会接着执行case2 输出ok2 !!只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

	//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的变量类型。

	var x interface{} //空接口、
	var y = 10.0
	x = y
	switch i := x.(type) { //x.(type)会显示x的真正数据类型
	case nil:
		fmt.Printf(" x的类型是 ：%T ", i)
	case int:
		fmt.Printf("x是int型")
	case float64:
		fmt.Printf("x 是float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或者 string型")
	default:
		fmt.Println("未知型")

	}

}
