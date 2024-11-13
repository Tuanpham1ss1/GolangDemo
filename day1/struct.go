package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	//basic struct
	p := Point{1, 2}
	fmt.Println(p, p.x, p.y)
	p.x = 3
	p.y = 4
	fmt.Println(p, p.x, p.y)
	//pointer to struct
	var a *Point = &p
	fmt.Println(a, *a)
	fmt.Println(a.x, (*a).x)
	//init struct
	var (
		v1 = Point{1, 2}
		v2 = Point{x: 1}
		v3 = Point{}
		k  = &Point{1, 2}
	)
	fmt.Println(v1, v2, v3, k)
}
