package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Printf("now=%v\n now type=%T", now, now)
}
