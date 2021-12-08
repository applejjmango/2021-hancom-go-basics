package main

import "fmt"

// Convert different types of value
func main() {
	gasPrice := 0.95859654305 // float64 type
	gasUsed  := 35000 // int //

	txFees := int(gasPrice * float64(gasUsed))
	fmt.Println(txFees)
}