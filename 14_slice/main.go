package main

import "fmt"

// Slice => Reference Type (복사가 이루어질 때 값 복사가 아닌, 참조 복사를 한다)
func main() {
	// Slice
	/*
	rainbow := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	rainbow = append(rainbow, "sky", "pink", "purple")
	fmt.Println(rainbow)
	*/

	// make 함수를 통한 slice 생성
	x := []int{2,4,6,8,10,12,14}
	fmt.Println(len(x)) // 길이 측정 
	fmt.Println(cap(x)) // 용량을 측정

	// []int 타입의 슬라이스는 len이 4개, cap이 6개
	slice := make([]int, 4, 6) // make(슬라이스 타입, 슬라이스 길이, 슬라이스 용량) 
	fmt.Println(slice) 

	slice = append(slice, 1,2,3,4,5)
	length := len(slice)
	cap := cap(slice)
	fmt.Println(slice, length, cap)
}



// func changeArr(arr [3]int) {
// 	arr[0] = 100 // 함수에 인자로 들어간 arr이라름 파라미터가 함수 스코프(중괄호) 내에서호지역변수로 생성, 즉, 원본의 메모리 주소 자체와 다름

// 	fmt.Println("This is Array =>", arr)
// }
// func changeSlice(slice []int) {
// 	slice[0] = 100 // 여기에서의 slice는 참조값이므로 원본 값 자체를 변경한다.
// 	fmt.Println("This is Slice Func => ", slice)
// }
// func main() {
// 	arr := [3]int{1,2,3} // [1,2,3]
// 	slice := []int{1,2,3,4}

// 	changeArr(arr)
// 	changeSlice(slice)

// 	fmt.Println("Main func array =>", arr) // 원본은 바뀌지 않음
// 	fmt.Println("Main func slice => ", slice) // 참조값이므로 원본이 바뀌었다.
// }