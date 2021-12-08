package main

import (
	"fmt"
	"strings"
)

// len, upper(대문자)
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("michel")
	fmt.Println(totalLength, uppercase)
}

