package main

import "fmt"

// naked return Func
func split(sum int) (x, y int) {
	x = sum * 5 / 4
	y = sum - x
	return
}

func main() {
	x, y := split(100)
	fmt.Println(x, y)
}
