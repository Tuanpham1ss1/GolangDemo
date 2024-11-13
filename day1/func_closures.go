package main

import "fmt"

func add1(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func subtract1(x float64) func(float64) float64 {
	return func(y float64) float64 {
		return x - y
	}
}

func main() {
	value := add1(3)(4)
	fmt.Println(value) // 7
	fmt.Println(subtract1(9.0)(8.12))
}
