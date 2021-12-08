package main

import "fmt"

// switch case
/*
func main() {
	switch day := 5; day {
		case 0:
			fmt.Println("Sunday")
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("TuesDay")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
	}
}*/

func canIDrive(age int) bool {
	switch {
	case age < 20:
		return false
	case age > 80:
		return false
	case age >= 20:
		return true
	}
	return false
}
func main() {
	fmt.Println(canIDrive(25))
}
