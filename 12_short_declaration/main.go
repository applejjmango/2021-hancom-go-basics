package main

import "fmt"

// // short declaration


func main() {
	width, height := 100, 50

  // DONT!
// width = 50 // assign 50 to width
// color := "red" // new variables
// fmt.Println(width, height, color)

//same as Above 👍
	width, color := 50, "Blue"
	fmt.Println(width, height, color)
}