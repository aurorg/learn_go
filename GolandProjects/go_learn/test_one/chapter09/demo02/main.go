package main

import "fmt"

func main() {
	cities := make(map[string]string)
	cities["1"] = "beijing"
	cities["2"] = "xian"
	cities["3"] = "nanjing"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	//遍历复杂的map
	studentMap := make(map[string]map[string]string)

	studentMap["01"] = make(map[string]string, 3)
	studentMap["01"]["name"] = "tom"
	studentMap["01"]["sex"] = "boy"
	studentMap["01"]["address"] = "xxx"

	studentMap["02"] = make(map[string]string, 3)
	studentMap["02"]["name"] = "mary"
	studentMap["02"]["sex"] = "girl"
	studentMap["02"]["address"] = "xxx"

	for k1, v1 := range studentMap {
		fmt.Println("k1", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v v2=%v\n", k2, v2)
		}
		fmt.Println()
	}

}
