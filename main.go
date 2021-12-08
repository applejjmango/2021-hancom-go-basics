package main

import (
	"fmt"
	"os"
)
func HelloWorld() {
	file, err := os.Open("hello.txt")
	defer file.Close() // defer 용법, 파일을 열었기 때문에, 작업이 끝나면 파일을 닫는다
	if err != nil {
		fmt.Println(err)
		return
	}
	buffer := make([]byte, 1024)
	fmt.Println("buffer =< ", buffer)
	if _, err = file.Read(buffer); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buffer))
}
// defer
// defer는 함수 내에서 제일 나중에, 즉 끝나기 직전에 실행하는 용법이다
func main() {
	HelloWorld()
	fmt.Println("Finished!!")
	/* 1 번
	defer fmt.Println("World")
	fmt.Println("Hello")
	*/

	/* 2 번
	defer func(){
		fmt.Println("defer is called at the last")
	}()
	fmt.Println("This is a main call")
	*/

    // 3번
	/*
	var num1, num2 int = 50, 0
	
	defer fmt.Println("Finished!!!")
	
	result := num1 / num2  
	fmt.Println(result)
	*/
}





/* make 함수를 통해서 slice 생성
// 
func main() {
	x := []int{2,4,6,8,10,12,14}
	fmt.Println(len(x)) // 길이 측정 
	fmt.Println(cap(x)) // 용량을 측정

	// []int 타입의 슬라이스는 len이 4개, cap이 6개
	slice := make([]int, 4, 6) // make(슬라이스 타입, 슬라이스 길이, 슬라이스 용량) 
	fmt.Println(slice) 

	slice = append(slice, 1,2,3)
	length := len(slice)
	cap := cap(slice)
	fmt.Println(slice, length, cap)
}
*/

/* Reference type => slice
func changeArr(arr [3]int) {
	arr[0] = 100 // 함수에 인자로 들어간 arr이라름 파라미터가 함수 스코프(중괄호) 내에서호지역변수로 생성, 즉, 원본의 메모리 주소 자체와 다름

	fmt.Println("This is Array =>", arr)
}
func changeSlice(slice []int) {
	slice[0] = 100 // 여기에서의 slice는 참조값이므로 원본 값 자체를 변경한다.
	fmt.Println("This is Slice Func => ", slice)
}
func main() {
	arr := [3]int{1,2,3} // [1,2,3]
	slice := []int{1,2,3,4}

	changeArr(arr)
	changeSlice(slice)

	fmt.Println("Main func array =>", arr) // 원본은 바뀌지 않음
	fmt.Println("Main func slice => ", slice) // 참조값이므로 원본이 바뀌었다.
}
*/

/*
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
*/

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
/*
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
*/

/* Struct
// struct
type User struct {
	name string
	occupation string
	age int
	hobbies []string
}

func main() {
	user := User{
		name: "Kim",
		occupation: "Technical Manager",
		age: 30,
		hobbies: []string{"soccer", "basketball"},
	}
	fmt.Printf("%s is %d years old and he is a %s\n", user.name, user.age, user.occupation)
}
*/

// Map
/*
func main () {
	myDog := map[string]string{"name": "doy", "gender": "male", "age": "3"}
	for key, value := range myDog {
		fmt.Println(key, value)
	}
}
*/

// Array (배열)
/*
func main() {
	// Array
	var rainbow [7]string // declaration

	rainbow[0] = "red"
	rainbow[1] = "orange"
	rainbow[2] = "yellow"
	rainbow[3] = "green"
	rainbow[4] = "blue"
	rainbow[5] = "indigo"
	rainbow[6] = "violet"

	// same as above
	rainbow2 := [7]string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}

	fmt.Println(rainbow)
	fmt.Println(rainbow2)
}
*/

/* Slice
// Slice
func main() {
	// Slice
	rainbow := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	rainbow = append(rainbow, "sky", "pink", "purple")
	fmt.Println(rainbow)
}
*/

// Convert different types of value
/*
func main() {
	gasPrice := 0.95859654305 // float64 type
	gasUsed  := 35000 // int //

	txFees := int(gasPrice * float64(gasUsed))
	fmt.Println(txFees)
}
*/

// // short declaration
// func main() {
// 	width, height := 100, 50
//  DONT!
// width = 50 // assign 50 to width
// color := "red" // new variables
// fmt.Println(width, height, color)
// same as Above 👍
// 	width, color := 50, "Blue"
// 	fmt.Println(width, height, color)
// }

// switch case
/*
func main() {
	switch day := 5; day {
		case 0:
			fmt.Println("Sunday")
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("TuesDay")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
	}
}*/
/*
func canIDrive(age int) bool {
	switch {
	case age < 20:
		return false
	case age > 80:
		return false
	case age >= 20:
		return true
	}
	return false
}
func main() {
	fmt.Println(canIDrive(25))
}
*/

/* 1번 // for, range
func add(numbers ...int ) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count
}
func main () {
	result := add(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(result)
}
*/
/* 2번
func add(numbers ...int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}
func main() {
	add(10,20,30,40,50)
}
*/

// public vs private
// zoo 폴더를 참고

/* len, upper(대문자)
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, uppercase := lenAndUpper("michel")
	fmt.Println(totalLength, uppercase)
}
*/

/* naked return Func
func split(sum int) (x, y int) {
	x = sum * 5 / 4
	y = sum - x
	return
}

func main() {
	x, y := split(100)
	fmt.Println(x, y)
}
*/

/* n개 이상의 params를 전달 받는 Func
func favouriteColor(color ...string) {
	fmt.Println(color)
}

func main () {
	favouriteColor("yellow", "blue", "red", "dark-grey", "black", "sky")
}
*/

/* unused return value
func koreanLegalDrinkingAge(age int) (int, string) {
	return age + 1, "This is Only for Koreans"
}
func main() {
	// age, words := koreanLegalDrinkingAge(30)
	// fmt.Println(age, words)
	age, _ := koreanLegalDrinkingAge(30)
	fmt.Println(age)
}
*/

/* function and types
func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func add3(a, b int) string {
	return "Kim " + "Young-hee"
}

func main() {
	fmt.Println(add3(10, 10))
}
*/

/* // Data Types
func main() {
	var i int = 500 // int => integer
	var u uint = uint(i) // unsigned
	var f float32 = float32(i)

	println(f, u)

	str := "ABCDEF"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
*/

// 03_variables_constants
// firstName4 := "Mary" // (X) 축약 문법은 오로지 함수 내에서 사용 가능
// var firstName4 string = "Mary" // (O)
/*
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
*/

// 02_comment
// func main() {
/*
	fmt.Println('저는 주석이라서 출력이 안되요')
*/
// fmt.Println('저는 한 줄 주석입니다')
// }

// 01_hello
/*
func main() {
	fmt.Println("Hello, ", "I'm Kim", "30 years old!")	// 쉼표(,)를 구분해 여러 인자를 전달해서 사용가능하다.

	fmt.Print("Hello ~, ", "I'm Jack ") // 인자마다 띄어쓰기를 하지 않고 마지막에 줄바꿈을 하지 않아요
	fmt.Printf("%d years old!", 50) // 형식 지정자를 사용할 수 있다.


}
*/

