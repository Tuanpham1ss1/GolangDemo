package main

import "fmt"

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}
	fmt.Print("done ")

	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)
}
