package main

import "fmt"

// n개 이상의 params를 전달 받는 Func
func favouriteColor(color ...string) {
	fmt.Println(color)
}

func main () {
	favouriteColor("yellow", "blue", "red", "dark-grey", "black", "sky")
}
