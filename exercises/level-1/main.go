package main

import "fmt"

func question_4() {
	// underlying type of mytype is string
	// https://golang.org/ref/spec#Types
	type mytype string
	var v mytype = "hello world"
	fmt.Printf("==============================\n")
	fmt.Printf("type = %T\nvalue = %v\n", v, v)
	fmt.Printf("==============================\n\n")
}

func question_5() {
	// underlying type of mytype is string
	// https://golang.org/ref/spec#Types
	type mytype string
	var v mytype = "hello world"

	// assigning value from mytype "v" to string type "x"
	// This is called as conversion and not casting
	var x string
	x = string(v)

	fmt.Printf("==============================\n")
	fmt.Printf("type = %T\nvalue = %v\n", v, v)
	fmt.Println(x)
	fmt.Printf("==============================\n\n")
}
func main() {
	question_4()
	question_5()
}
