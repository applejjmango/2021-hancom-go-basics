package main

import "fmt"

// Pointer (포인터) => 어떤 값의 메모리 주소를 의미한다.
/* 1번
func main() {
	a := 100
	b := 200

	fmt.Println(&a, &b) // 각각의 a와 b의 메모리 주소를 가리킨다.

	c := 300
	d := &c
	fmt.Println("c, d =>", &c, d) // 2개 모두 c의 메모리 주소 값이 나오는 것을 확인할 수 있다.
	fmt.Println("C의 값 =>", *d) // c(메모리 주소)를 통해서 c의 값인 300을 출력하게 된다

	c = 500
	fmt.Println(*d) // c의 값이 변경이 된 상태에서, d는 c의 메모리 주소를 바라보고 있기 때문에 결과적으로 c의 변경된 값을 출력한다.

	e := 1000
	f := &e
	fmt.Println("f =>", f)
	*f = 2000
	fmt.Println("e => ", e) // f가 e의 메모리 주소를 바라 보고 있는 상태에서, *f를 통해서 e의 실제 값을 변경했다.
}
*/

func main() {
	i, j := 50, 100
	fmt.Println("1", &i, &j) // 메모리 주소  

	p := &i
	fmt.Println("2", p) 
	fmt.Println("3", *p)

	*p = 20
	fmt.Println("4", i)

	*p = *p / 5
	fmt.Println("5", j)
	fmt.Println("6", p)
}
