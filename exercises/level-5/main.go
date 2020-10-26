package main

import "fmt"

// Documentation on struct from golang lang spec
// spec => struct types
// named elements each of which has a name and type.
// we can have an anonymous field e.g. person in below e.g.
// an empty struct e.g. `struct {}`
// struct {
// 	x, y int
// 	u float32
// 	_ float32 // padding
// 	A *[]int
// 	F func()
// }
// trick for swapping
// x, y = y, x
// A field declared with a type but no explicit field name is an anonumous
// field, also called an embedded field or an embeddeding of the type in the
// struct. An embedded type must be specified as a type name T or as a pointer
// to a non-interface type name *T, and T may not be a pointer type.
// ** The unqualified type name acts as the fielf name. **
// that's why it is accessed as sa2.person.first

// Goals of the language ;)
// efficient compilation, efficient execution and ease, ease of programming.

func question_1() {

	fmt.Printf("\n\n============= Q-1 ================\n")
	type person struct {
		first string
		last  string
		age   int
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Println(p1)

	p2 := person{
		first: "1",
		last:  "11",
		age:   111,
	}
	var p []person = []person{p1, p2}
	// or we can declare it like `p := []person{p1, p2}`
	fmt.Println(p2)
	for _, v := range p {
		fmt.Println(v)
	}

	// o/p of above code is like below
	// ============= Q-1 ================
	// {James Bond 32}
	// {1 11 111}
	// {James Bond 32}
	// {1 11 111}
	// =============================
	fmt.Printf("=============================\n")
}

func question_2() {
	fmt.Printf("\n\n============= Q-2 ================\n")
	type person struct {
		first string
		last  string
		age   int
	}
	p1 := person{
		first: "1",
		last:  "11",
		age:   111,
	}
	p2 := person{
		first: "2",
		last:  "22",
		age:   222,
	}
	fmt.Println(p1, p2)
	mp := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	fmt.Println(mp)
	fmt.Printf("=============================\n")

	// o/p is like below
	// ============= Q-2 ================
	// {1 11 111} {2 22 222}
	// map[1:{1 11 111} 2:{2 22 222}]
	// =============================
	//
	// below prints like below =
	//
	// 2
	// 2
	// 22
	// 222
	// 1
	// 1
	// 11
	// 111
	// =============================
	for k, v := range mp {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		fmt.Println(v.age)
	}
	fmt.Printf("=============================\n")
}

func question_3() {
	fmt.Printf("\n\n============= Q-3 ================\n")
	type vehicle struct {
		doors int
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{ // vehicle : vehicle {} is very imp.
			doors: 2,
			color: "brown",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Printf("=============================\n")
}

// create an anonymous struct and use it.
func question_4() {

	fmt.Printf("\n\n============= Q-4 ================\n")
	anon := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"water",
		},
	}
	fmt.Println(anon)
	fmt.Printf("=============================\n")

	// o/p is like below :=
	// ============= Q-4 ================
	// {James map[M:888 Moneypenny:555 Q:777] [Martini water]}
	// =============================

	// some more e.g.
	for k, v := range anon.friends {
		fmt.Println(k, v)
	}

	for i, v := range anon.favDrinks {
		fmt.Println(i, v)
	}

	// o/p of above is like
	// Moneypenny 555
	// Q 777
	// M 888
	// 0 Martini
	// 1 water
}

func main() {
	// struct aggregates together different data types.
	// this is user defined type
	// struct is a composite data type aka aggregate data type
	// allows us to compose together values of different types.
	type person struct {
		first string
		last  string
	}

	p1 := person{
		first: "James",
		last:  "Bond",
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first)
	fmt.Println(p2.last)

	// embedded types
	type secretAgent struct {
		person
		ltk bool
	}

	// note here that the embedded type fields from person got promoted
	// to struct secretAgent. So we need not use sa1.person.first
	// we can directly use sa1.first. Although both are valid and in fact
	// in case of name collision we will use sa1.person.first itself.
	// innertype just got promoted to outer type.
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1.first, sa1.last, sa1.ltk)
	fmt.Println(sa1.person.first, sa1.first)

	// name collision e.g.
	type secretArmy struct {
		person
		first string
		last  string
		ltk   bool
	}

	sa2 := secretArmy{
		person: person{
			first: "James",
			last:  "Bond",
		},
		first: "CIA",
		last:  "CIA2",
		ltk:   true,
	}
	fmt.Println(sa2.person.first, sa2.first)
	// o/p till above :=
	// {James Bond}
	// {Miss Moneypenny}
	// James
	// Moneypenny
	// James Bond true
	// James James
	// James CIA

	// Anonymous struct
	type person2 struct {
		first string
		last  string
		age   int
	}
	p3 := person2{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	fmt.Printf("%T, %v\n", p3, p3) // main.person2, {James Bond 32}

	// so instead of doing above we could do composite literal without the
	// name of the struct and called anonymous struct.
	// it is used for reducing code pollution and it's use is limited.
	// To keep the code lean and clean.
	p4 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   33,
	}
	fmt.Printf("%T %v\n", p4, p4) // struct { first string; last string; age int } {James Bond 33}

	question_1()
	question_2()
	question_3()
	question_4()

}
