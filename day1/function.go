package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}
func swap(a, b int) (int, int) {
	return b, a
}
func split(number int) (x, y int) {
	x = number / 10
	y = number % 10
	return
}
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(subtract(1, 2))
	a, b := swap(1, 2)
	fmt.Println(a, b)
	fmt.Println(split(12))
}
