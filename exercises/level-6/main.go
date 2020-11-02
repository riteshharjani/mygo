package main

import (
	"fmt"
	"math"
)

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

	// anonymous function : are functions with no function name
	// e.g. (remember func is a reserved keyword for function types
	// func(<args>) <return> { <...code...>}()
	// functions are first class citizens. Sicne function is a type
	// like any other type. e.g. of type int, string etc.
	// so function is a first class citizen.
	// hence we can assign function to variables, pass in as argument
	// and return it as well
	// e.g.
	// v := func() {}
	// v()
	// Returning a func()
	func() {
		fmt.Println("I am an anonymous function, foo")
	}()
	func(x int) {
		fmt.Println("I am an anonymous function, foo", x)
	}(42)

	v := func(x int) {
		fmt.Println("I am a function expression, foo", x)
	}
	v(43)

	// here a func expression which takes int x and returns a func(y int)
	// which should return an int.
	// and then inside 1st func expression, we are retuning the func
	// expression.
	s := func(x int) func(y int) int {
		fmt.Println("retuning a function bar()", x)
		return func(y int) int {
			fmt.Println("I am returned anon function", x, y)
			return y
		}
	}(42)
	fmt.Printf("type of s = %T\n", s)
	i := s(43)
	fmt.Println(i)
	// o/p of above code.
	// I am an anonymous function, foo
	// I am an anonymous function, foo 42
	// I am a function expression, foo 43
	// retuning a function bar() 42
	// type of s = func(int) int
	// I am returned anon function 42 43
	// 43

	// Callbacks - when we are passing a function as an argument
	// and then calling that callback function
	//
	// since func are first class citizen so it can be passed as an arg.
	// also called as functional programming.
	// below is a e.g. of function which do the sum of values
	// passed in as variadic parameter (which could be any number of params)
	// e.g. xi will be a slice if a slice is passed

	sm := func(xi ...int) int {
		tot := 0
		for _, v := range xi {
			tot += v
		}
		return tot
	}
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("====> callback", sm(ii...))

	// e.g. here we are defining a func ev which is taking
	// a function f as an argument which is a callback function
	// to do the sum (which is sm)
	// Now this function ev ranges over vi and creates another
	// slice yi which is slice of only even numbers which is passed
	// to a function sum to return the sum
	ev := func(f func(xi ...int) int, vi ...int) int {
		var yi []int
		for _, v := range vi {
			if v%2 == 0 {
				yi = append(yi, v)
			}
		}
		return f(yi...)
	}
	fmt.Println("====> callback func", ev(sm, ii...))
	// o/p
	// ====> callback 45
	// ====> callback func 20

	// Closure - scope of a variable.
	// clousre allows us to limit the scope of variables.
	// e.g. here y is not accessible from outside of scope
	// in 2nd e.g. if you see a := incrementor()
	// creates another scoped variable x inside scope of memory of a
	// so every time a() is called the incrementor() increases
	{
		y := 42
		fmt.Println("Hello from clousre", y)
	}
	a := incrementor()
	fmt.Println("value of a = ", a())
	fmt.Println("value of a = ", a())
	b := incrementor()
	fmt.Println("value of b = ", b())
	fmt.Println("value of b = ", b())

	// e.g. of above code o/p
	// Hello from clousre 42
	// value of a =  1
	// value of a =  2
	// value of b =  1
	// value of b =  2
	//
	question_1()
	question_2()
	question_3()
	question_4()
	question_5()
	question_6()
	question_7()
	question_8()
	question_9()
	question_10()
}

func incrementor() func() int {
	var x int = 0
	return func() int {
		x++
		return x
	}
}

// Exercises from here
// ===================
// Hands on exercise
// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results
func question_1() {
	// instead of creating a func with identifier foo
	// created a variable of type functioned named foo

	fmt.Println("================ Q-1 =================")
	foo := func() int {
		return 1
	}
	fmt.Println("foo returned", foo())

	bar := func() (int, string) {
		return 1, "hello world"
	}
	i, s := bar()
	fmt.Println("bar returned", i, s)
	fmt.Println("======================================")
	// ================ Q-1 =================
	// foo returned 1
	// bar returned 1 hello world
	// ======================================

}

// create a func with the identifier foo that
// 	takes in a variadic parameter of type int
// 	pass in a value of type []int into your func (unfurl the []int)
// 	returns the sum of all values of type int passed in
//
// create a func with the identifier bar that
// 	takes in a parameter of type []int
// 	returns the sum of all values of type int passed in
func question_2() {
	fmt.Println("=================Q-2=====================")
	foo := func(vi ...int) int {
		tot := 0
		for _, v := range vi {
			tot += v
		}
		return tot
	}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("sum of values = ", foo(arr...))

	bar := func(vi []int) int {
		tot := 0
		for _, v := range vi {
			tot += v
		}
		return tot
	}
	fmt.Println("sum of values = ", bar(arr))
	fmt.Println("======================================")
	// =================Q-2=====================
	// sum of values =  45
	// sum of values =  45
	// ======================================
}

// use the defer keyword to show that a deferred func runs after the
// function exits
func question_3() {
	fmt.Println("=================Q-3=====================")
	x := 1
	defer func() {
		x = 3
		fmt.Println("called from defer, value should be 3", x)
		fmt.Println("======================================")
	}()
	x = 5
	fmt.Println("the value of x should be 5", x)
	// =================Q-3=====================
	// the value of x should be 5 5
	// called from defer, value should be 3 3
	// ======================================
}

// Create a user defined struct with
// 	the identifier “person”
// 	the fields:
// 		first
// 		last
// 		age
// attach a method to type person with
// 	the identifier “speak”
// 	the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person
type Person struct {
	first string
	last  string
	age   int
}

func (x Person) speak2() {
	fmt.Println("My name is", x.first, x.last)
}

func question_4() {
	fmt.Println("=================Q-4====================")

	p := Person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p.speak2()
	fmt.Println("======================================")

}

// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct {
	len float64
}
type circle struct {
	rad float64
}

func (s square) area() float64 {
	return s.len * s.len
}

func (c circle) area() float64 {
	return math.Pi * c.rad * c.rad
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area from interface = ", s.area())
}

func question_5() {
	fmt.Println("=================Q-5=====================")
	s := square{
		len: 10,
	}
	c := circle{
		rad: 10,
	}
	fmt.Printf("area of square of len %f is %f\n", s.len, s.area())
	fmt.Printf("area of circle of len %f is %f\n", c.rad, c.area())
	info(s)
	info(c)
	fmt.Println("======================================")
	// =================Q-5=====================
	// area of square of len 10.000000 is 100.000000
	// area of circle of len 10.000000 is 314.159265
	// area from interface =  100
	// area from interface =  314.1592653589793
	// ======================================

}

// build and use an anonymous func
func question_6() {
	fmt.Println("================Q-6======================")
	func() {
		fmt.Println("I am an anon func")
	}()
	fmt.Println("======================================")
	// ================Q-6======================
	// I am an anon func
	// ======================================
}

// this is legal
var ff = func() int {
	fmt.Println("here is how you declare me outside in global scope")
	return 5
}

// compiler will complain for this
// var ff
// assign a func to a variable
func question_7() {
	fmt.Println("===============Q-7=======================")
	var f func(i int) int
	f = func(i int) int {
		return i * 10
	}
	fmt.Printf("return from f = %d\n", f(5))
	ff()
	fmt.Println("======================================")

	// ===============Q-7=======================
	// return from f = 50
	// here is how you declare me outside in global scope
	// ======================================
}

// defines a func which returns a func and then call that
func question_8() {
	f := func() func() int {
		return func() int {
			return 42
		}
	}
	fmt.Println(f()())
	// 42
}

// pass a func into a func as an argument and then call it
func question_9() {
	f := func(f func(i int) float64) {
		fmt.Println(f(42))
	}
	g := func(i int) float64 {
		return 3.14 * float64(i)
	}
	f(g)
	// 131.88
}

// show a use of closure
// here val could show the use of closure
func question_10() {
	f := func() func() int {
		val := 0
		return func() int {
			val++
			fmt.Println(val)
			return val
		}
	}()
	f()
	f()
	f()
	// 1
	// 2
	// 3
	//

}
