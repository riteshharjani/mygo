package main

import "fmt"

// func (r receiver) identifier (parameters) (return) { ... }
// everything in go is PASS by VALUE
//

// e.g. function
func foo() {
	fmt.Printf("hello from foo\n")
}

func bar(s string) {
	fmt.Printf("%s\n", s)
}

func woo(s string) string {
	return fmt.Sprintf("hello from woo, %s\n", s)
}

// variadic parameter function e.g.
// this takes an unlimited number of paramters as a slice argument
// it has to be the final parameter. If nothing is paased it will consider
// as a nil slice
// Otherwise a new underlying array is creasted and passed a new slice is passed
// of type []T.
// But if the final argument is assignable to a slice type []T, it may be passed
// unchanged as the value for a ...T param if the argument is followed by ...
// In this case no new slice is created.
//
// Basically if we are passing a slice by unfurling method, then no new slice
// is created and the slice passed will have the same value as x with the underlying
// array. And in case if you change the value of the array in function sum
// then the value of actual array will also be changed in the callee function.
//
// this is a variadic paramter of any type.
// int is also of type empty interface.
// all types are of type empty interface.
// func Println(a ...interface{})
func sum(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// e.g. output, till above.
	// [1 2 3 4 5 6]
	// []int

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("total is ", sum)
	if len(x) != 0 {
		x[0] = 100
	}
}

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

//func (r receiver) identifier (arguments) (return(s)) { code }
// when you have a receiver it will attach this function `speak` to
// `type secretAgent struct`. Then it has access to the receiver
func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "- from human")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "- from person")
}

// now we will learn about interfaces. It allow us to define behaviors and also
// allows us to do polymorphism
// keyword identifier type (we follow the same analogy for all the types)
// e.g. var x int (keyword is var, identifier is x and type is int)
// Now a value can be of more than one type.
// so any other type that has the function speak() is also of the type human
// a value can be of more than one type. (interfaces allows this)
// so here sa1 (value) is of type (secretAgent) assigned to identifier sa1.
// but because it has method speak() attached to it so it is also of type
// human.
//
// e.g. hence
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }
// hence any type that also defines a method Write is also of type
// Writer or say it implements the Writer interface.
// So with this Writer interface allows us to pass any type in the function
// argument that implements a Write method.
// and all those different type can implement this method a bit differently.
//
// an interface says :- hey baby!, if your type got this method then you are my type.
// ;)
//
type human interface {
	speak()
}

func bar_(h human) {
	switch h.(type) { // this is asserting
	case person:
		fmt.Println(" I was passed into barr switch person", h.(person).first) // this is also asserting that this is of type person.
		// so h.(person) will be saying h.(person) will be of type person
	case secretAgent:
		fmt.Println("I was passed into bar switch secretAgent", h.(secretAgent).person.first, h.(secretAgent).first) // see previous struct execises
		// note this is needed since we cannot access h.first directly.
		// So h.(secretAgent) assert that I am of type secretAgent and so go an access .first field
	}
	fmt.Println("I was called in bar", h)
}

func main() {
	foo()
	bar("hello")
	fmt.Println(woo("Miss"))
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	// unfurling
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//this below gives error = ./main.go:47:5: cannot use x (type []int) as type int in argument to sum
	//sum(x)

	// below is called unfurling
	sum(x...)
	fmt.Println(x) // [100 2 3 4 5 6 7 8 9] // x gets changed
	//
	// variadic means 0 or more.
	// so below is legal
	sum() // total is  0

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			32,
		},
		ltk: true,
	}
	sa1.speak() // output "I am  James Bond"

	p := person{
		"1",
		"11",
		111,
	}
	p.speak() // I am  1 11 - from person

	// here is an example of polymorphism
	// bar is taking a secretAgent and person.
	// This happens since they are also of type human.
	// So this is called polymorphism.
	// interfaces allows value to be of many different types.
	bar_(sa1) // I was called in bar {{James Bond 32} true}
	bar_(p)   // I was called in bar {1 11 111}

	// conversion e.g.
	type hotdog int
	var xx hotdog = 42
	var yy int = int(xx)
	fmt.Printf("%T %v\n", xx, xx) // main.hotdog 42
	fmt.Printf("%T %v\n", yy, yy) // int 42

	//assertion
	/*
		switch xx.(type) { // this is not valid. It says cannot apply switch on non-interface value xx. (type hotdog)
		case int:
			fmt.Println("I am int", xx)
		case hotdog:
			fmt.Println("I am hotdog", xx)
		}
	*/
}
