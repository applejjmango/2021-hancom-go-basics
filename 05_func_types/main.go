package main

import "fmt"

// function and types
func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func add3(a, b int) string {
	return "Kim " + "Young-hee"
}

func main() {
	fmt.Println(add3(10, 10))
}
