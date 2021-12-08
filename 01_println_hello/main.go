// 01_hello
package main

import "fmt"

func main() {
	fmt.Println("Hello, ", "I'm Kim", "30 years old!")	// 쉼표(,)를 구분해 여러 인자를 전달해서 사용가능하다.

	fmt.Print("Hello ~, ", "I'm Jack ") // 인자마다 띄어쓰기를 하지 않고 마지막에 줄바꿈을 하지 않아요
	fmt.Printf("%d years old!", 50) // 형식 지정자를 사용할 수 있다.
}