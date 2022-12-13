package main

import "fmt"

func main() {
	//切片的数据类型为map
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "aaa"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "bbb"
		monsters[1]["age"] = "400"
	}
	//前面定义了两个，多了的话就会出现越界
	//这时就可以使用切片的append函数，可以动态的增
	newMonster := map[string]string{
		"name": "333",
		"age":  "222",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}

