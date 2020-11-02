package main

import "fmt"

// I see everything above pointers in go is same as C language
// So I didn't see any other difference
//
// Note - everything in go is pass by value
// are you passing a value or an address. Just everything in go is pass
// by value

func foo(x int) {
	fmt.Println(&x)
	fmt.Println(x)
	x = 43
	fmt.Println(&x)
	fmt.Println(x)
}

func _foo(y *int) {
	fmt.Println(y)
	fmt.Println(*y)
	*y = 43
	fmt.Println(y)
	fmt.Println(*y)
}

func main() {
	fmt.Println("Pointers example")
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	// or b := &a
	fmt.Println(b)
	// note "*" is an operator
	fmt.Println(*b, *&a)

	// o/p of above
	// Pointers example
	// 42
	// 0xc000014110
	// int
	// *int
	// 0xc000014110
	// 42 42
	//

	// this shows x is passed by value
	x := 0
	fmt.Println("===>", x)
	fmt.Println(&x)
	foo(x)
	fmt.Println(x)
	fmt.Println(&x)
	// ===> 0
	// 0xc000014138
	// 0xc000014140
	// 0
	// 0xc000014140
	// 43
	// 0
	// 0xc000014138

	y := 1
	fmt.Println("===>", y)
	fmt.Println(&y)
	_foo(&y)
	fmt.Println(y)
	fmt.Println(&y)
	// ===> 1
	// 0xc0000b2060
	// 0xc0000b2060
	// 1
	// 0xc0000b2060
	// 43
	// 43
	// 0xc0000b2060

	// method sets
	// what are the types attached to methods, those are called method sets
	// determins what methods are attached to a given TYPE.
	// It is exactly what it says - what is the set of methods for a given
	// TYPE? That is its method set.
	// a NON POINTER RECIEVER
	//  =    works with values which are pointers or non-pointers.
	// a POINTER RECIEVER
	//  =    only works with values that are POINTERS
	//
	//  RECEIVERS 		VALUES
	//  (t T) 		T and *T
	//  (t *T) 		*T
	// Method sets
	// A type may have a method set associated with it. The method set of an interface type is its interface. The method set of any other type T consists of all methods declared with receiver type T. The method set of the corresponding pointer type *T is the set of all methods declared with receiver *T or T (that is, it also contains the method set of T). Further rules apply to structs containing embedded fields, as described in the section on struct types. Any other type has an empty method set. In a method set, each method must have a unique non-blank method name.
	//
	// The method set of a type determines the interfaces that the type implements and the methods that can be called using a receiver of that type.

	c := circle{
		rad: 10.2,
	}
	// notice the reciever type can be a pointer or a non-pointer
	// it can be called from either type (pointer value) or non-pointer value. Go compiler takes care of the conversion for us
	// only if the method set is declared with a pointer type only then the value of the struct can be changed
	fmt.Println("area:", c.area(), c.rad)
	c.change_rad(2.1)
	fmt.Println("area:", c.area(), c.rad)
	// o/p of above code
	// area: 13.8474 10.2
	// area: 13.8474 2.1

}

type circle struct {
	rad float64
}

func (c circle) area() float64 {
	c.rad = 2.1
	return 3.14 * c.rad * c.rad
}

func (c *circle) change_rad(r float64) {
	c.rad = r
}
