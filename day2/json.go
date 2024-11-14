package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct {
	Name, Color string
	Age         int
	Weight      float64
	IsMale      bool
}

func main() {
	a := Cat{
		Name:   "Tuan",
		Color:  "yellow",
		Age:    18,
		Weight: 75,
		IsMale: true,
	}
	b := Cat{
		Name:   "Pham",
		Color:  "black",
		Age:    23,
		Weight: 87,
		IsMale: false,
	}
	aJson, err := json.Marshal(a)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(string(aJson))
	bJson, _ := json.MarshalIndent(b, "", "  ")
	fmt.Println(string(bJson))
	jsonData := []byte(`{"Name":"Tuan","Color":"yellow","Age":18,"Weight":75,"IsMale":true}`)
	isValid := json.Valid(jsonData)
	fmt.Println(isValid)
	var chiha Cat
	err = json.Unmarshal(jsonData, &chiha)
	if err != nil {
		println(err)
		return
	}
	fmt.Println(chiha)
}
