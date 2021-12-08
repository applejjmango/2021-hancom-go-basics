package main

import "fmt"

// 03_variables_constants
// firstName4 := "Mary" // (X) 축약 문법은 오로지 함수 내에서 사용 가능
// var firstName4 string = "Mary" // (O)

func main() {
    // const
	const firstName string = "Mary"
	// firstName = "John"
	fmt.Println(firstName)

	// var
	var firstName2 string = "Mary"
	firstName2 = "John"
	fmt.Println(firstName2)

	// 축약 문법
	firstName3 := "Mary" // var firstName2 string = "Mary 와 같다
	firstName3 = "John"
	age := 30
	fmt.Println(firstName3, age)
}
