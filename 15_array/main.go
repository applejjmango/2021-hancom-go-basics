package main

import "fmt"

// Array (배열) => Value Type

func main() {
	// Array
	var rainbow [7]string // declaration
	

	rainbow[0] = "red"
	rainbow[1] = "orange"
	rainbow[2] = "yellow"
	rainbow[3] = "green"
	rainbow[4] = "blue"
	rainbow[5] = "indigo"
	rainbow[6] = "violet"

	// same as above
	rainbow2 := [7]string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}

	fmt.Println(rainbow)
	fmt.Println(rainbow2)
}
