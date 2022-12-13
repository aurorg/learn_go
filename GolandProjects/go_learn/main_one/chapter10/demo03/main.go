package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1.LeftUp.x 地址=%p\n r1.leftUp.y 地址=%p\n r1.rightDown.x 地址=%p\n  r1.rightDown.y地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	fmt.Println()

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp 本身地址=%p     r2.rightDown 本身地址=%p \n",
		&r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp 指向的地址=%p    r2.rightDown 指向的地址=%p \n",
		r2.leftUp, r2.rightDown)
}
