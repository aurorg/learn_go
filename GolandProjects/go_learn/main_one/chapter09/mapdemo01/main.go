package main

import "fmt"

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["1"] = "111"
	a["2"] = "222"
	fmt.Println(a)

	cities := make(map[string]string)
	cities["1"] = "2323"
	cities["2"] = "dhyfahf"
	cities["3"] = "asdaf"
	fmt.Println(cities)

	heroes := map[string]string{
		"h1": "aaa",
		"h2": "bbb",
		"h3": "ccc",
	}
	heroes["h4"] = "sadad"
	fmt.Println("heroes=", heroes)

	delete(cities, "1")
	fmt.Println(cities)
}
