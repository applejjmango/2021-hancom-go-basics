package main

import "fmt"

// unused return value

func koreanLegalDrinkingAge(age int) (int, string) {
	return age + 1, "This is Only for Koreans"
}
func main() {
	// age, words := koreanLegalDrinkingAge(30)
	// fmt.Println(age, words)
	age, _ := koreanLegalDrinkingAge(30)
	fmt.Println(age)
}
