package main

import "fmt"

func question_4() {
	// underlying type of mytype is string
	// https://golang.org/ref/spec#Types
	type mytype string
	var v mytype = "hello world"
	fmt.Printf("type = %T\nvalue = %v\n", v, v)
}

func main() {
	question_4()
}
