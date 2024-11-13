package main

import "fmt"

var car bool
var age int = 21

func main() {
	var name string = "Tuan Pham"
	var job, address = "Software Engineer", "Hanoi"
	gender := "male"
	const homeTown = "Hanoi"
	fmt.Println(name, job, address, gender, car, age, homeTown)
	//basic type
	var (
		name1 string = "Tuan Pham"
		age1  int    = 23
	)
	fmt.Printf("type: %T, value: %v\n", name1, name1)
	fmt.Printf("type: %T, value: %v\n", age1, age1)
}
