package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello wolrd")
}
