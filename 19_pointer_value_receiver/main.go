package main

import "fmt"

// Method (메서드) => 함수명 앞에 타입과 변수명이 붙어 있는 것을 리시버(Receiver)라고 한다. func과 비슷한 모양을 띈다.
type Number struct {
	A, B int
}

// 메서드 (1. value receiver) => Number라고 정의된 struct에 아무것도 붙여 있지 않는 것은 벨류 리시버 라고 한다.
func (n Number) printNumber() {
	fmt.Println(n.A)
	fmt.Println(n.B)
}

// 메서드 (pointer receiver) 
// 값이 복제되는 것이 아니라, n이라는 원본 그 자체를 의미한다. (포인터로 가리키고 있기 때문에)
func (n *Number) addPointer(_number int) {
	n.A += _number
	n.B += _number
}

// 메서드 (value receiver) => n 자체가 복제된 값을 의미한다.다
// 즉, 완전히 다른 값을 의미하기 때문에, n을 출력을 하더라도 해당 부분의 결과 값이 반영이 되지 않는다
func (n Number) addValue(_number int) {
	n.A += _number
	n.B += _number
}

func main() {
	// n := Number{10, 20, 30, 40}
	// n.printNumber()
	n := Number{10, 20}
	fmt.Println("1", n)
	n.addPointer(100) 
	n.addValue(200)
	fmt.Println("2", n)
	
}