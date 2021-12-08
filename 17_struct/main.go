package main

import "fmt"

// struct
type User struct {
	name string
	occupation string
	age int
	hobbies []string
}

func main() {
	user := User{
		name: "Kim",
		occupation: "Technical Manager",
		age: 30,
		hobbies: []string{"soccer", "basketball"},
	}
	fmt.Printf("%s is %d years old and he is a %s\n", user.name, user.age, user.occupation)
}
