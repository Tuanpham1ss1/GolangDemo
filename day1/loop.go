package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}
	//while loop
	sum := 0
	for sum < 100 {
		sum += 10
	}
	fmt.Println(sum)
	//infinite loop
	for {
		fmt.Println("go")
	}
	//for range
	nums := []int{2, 3, 4}
	sum1 := 0
	for _, num := range nums {
		sum1 += num
	}
	fmt.Println("sum:", sum1)
	//map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

}
