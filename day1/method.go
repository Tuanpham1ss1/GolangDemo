package main

type User struct {
	name string
	age  int
}

func (u User) Greet() string {
	return "Hello, my name is " + u.name
}
func (u User) changename(name string) {
	u.name = name
}

func main() {
	u := User{"John", 25}
	u.changename("Doe")
	println(u.Greet())
}
