package main

import "fmt"

func main() {
	var list [7]int
	list[0] = 7
	fmt.Println(list, len(list))
	for i, v := range list {
		fmt.Println(i, v)
	}
	oddNumbers := [5]int{1, 3, 5, 7, 9}
	oddNumbers[2] = 11
	fmt.Println(oddNumbers)
	var fullName [4]string = [4]string{"John", "Doe", "is", "awesome"}
	fmt.Println(fullName)
}
