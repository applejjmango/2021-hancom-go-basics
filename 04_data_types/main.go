package main

// Data Types
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

