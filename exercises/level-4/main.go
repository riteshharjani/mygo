package main

import "fmt"

//func question_xxx() {
//	fmt.Printf("\n\n============= Q-xxx ==============\n")
//	fmt.Printf("===========================\n")
//}

func question_1() {

	/*
		Hands-on exercise #1
		Using a COMPOSITE LITERAL:
		create an ARRAY which holds 5 VALUES of TYPE int
		assign VALUES to each index position.
		Range over the array and print the values out.
		Using format printing
		print out the TYPE of the array
	*/

	fmt.Printf("\n\n============= Q-1 ==============\n")
	vi := [5]int{1, 2, 4, 5}
	for i, v := range vi {
		fmt.Println(i, v)
	}
	fmt.Printf("type of array = %T\n", vi)
	fmt.Printf("===========================\n")
}

func question_2() {
	/*
		Hands-on exercise #2
		Using a COMPOSITE LITERAL:
		create a SLICE of TYPE int
		assign 10 VALUES
		Range over the slice and print the values out.
		Using format printing
		print out the TYPE of the slice
	*/

	// vi := []int{}  // this will cause runtime error since length is 0
	// hence range overflow error
	// panic: runtime error: index out of range [0] with length 0

	fmt.Printf("\n\n============= Q-2 ==============\n")
	vi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//vi := make([]int, 10)

	// below overrides this values to a different value
	for i := 0; i < 10; i++ {
		vi[i] = i * 10
	}

	for i, v := range vi {
		fmt.Println(i, v)
	}
	fmt.Printf("===========================\n")
}

// slicing a slice
// Using the code from the previous example, use SLICING to create the
// following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
// [42 43 47 48 49]
func question_3() {
	fmt.Printf("\n\n============= Q-3 ==============\n")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(append(x[0:2], x[5:8]...)) // important

	fmt.Printf("===========================\n")
}

//Follow these steps:
// start with this slice
// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// append to that slice this value
// 52
// print out the slice
// in ONE STATEMENT append to that slice these values
// 53
// 54
// 55
// print out the slice
// append to the slice this slice
// y := []int{56, 57, 58, 59, 60}
// append is a builtin function that takes variadic number of arguments.
func question_4() {
	fmt.Printf("\n\n============= Q-4 ==============\n")

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println(x)

	fmt.Printf("===========================\n")
}

// To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
// start with this slice
// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
// [42, 43, 44, 48, 49, 50, 51]
func question_5() {
	fmt.Printf("\n\n============= Q-5 ==============\n")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
	fmt.Printf("===========================\n")
}

func question_6() {
	fmt.Printf("\n\n============= Q-6 ==============\n")

	x := make([]string, 50, 50)
	fmt.Println(x, len(x), cap(x))

	// this below assigning I guess will make a different slice all together
	// and assign it to x. It is clear from the output too.
	// ============= Q-6 ==============
	// [                                                 ] 50 50
	// [hello world this is a great golang course so far] 9 9
	// ===========================
	x = []string{"hello", "world", "this", "is", "a", "great", "golang", "course", "so far"}
	fmt.Println(x, len(x), cap(x))

	// for iterating
	for i := 0; i < len(x); i++ {
		fmt.Printf("%d %s\n", i, x[i])
	}
	fmt.Printf("===========================\n")
}

// Create a slice of a slice of string ([][]string).
// Store the following data in the multi-dimensional slice:
// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// o/p =
// ============= Q-7 ==============
// [[James Bond Shaken, not stirred] [Miss Moneypenny Helloooo, James.]]
// [[James Bond Shaken, not stirred] [Miss Moneypenny Helloooo, James.]]
// 0 => James
// 1 => Bond
// 2 => Shaken, not stirred
// 0 => Miss
// 1 => Moneypenny
// 2 => Helloooo, James.
// ===========================
func question_7() {
	fmt.Printf("\n\n============= Q-7 ==============\n")

	x := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooo, James."},
	}
	fmt.Println(x)

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooo, James."}
	xx := [][]string{xs1, xs2}
	fmt.Println(xx)

	//iterating over 2-d slices
	for _, xs := range xx {
		for j, v := range xs {
			fmt.Printf("%d => %v\n", j, v)
		}
	}
	fmt.Printf("===========================\n")
}

// Create a map with a key of TYPE string which is a person’s “last_first” name,
// and a value of TYPE []string which stores their favorite things. Store three
// records in your map. Print out all of the values, along with their index
// position in the slice.
// `bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
// `moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
// `no_dr`, `Being evil`, `Ice cream`, `Sunsets`
//
// o/p =
// ============= Q-8 ==============
// bond_james [Shaken, not stirred Martinis Women]
// moneypenny_miss [James Bond Literature Computer Science]
// fleming_ian [steaks cigars espionage]
// no_dr [Being evil ice cream Sunsets]
// ===========================
func question_8_9() {
	fmt.Printf("\n\n============= Q-8-9 ==============\n")

	mp := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "ice cream", "Sunsets"},
	}

	mp["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range mp {
		fmt.Println(k, v)
	}
	fmt.Printf("===========================\n")
}

// Using the code from the previous example, delete a record from your map.
// Now print the map out using the “range” loop
// solution: https://play.golang.org/p/TYl5EbjoeC
// o/p :=
// ============= Q-10 ==============
// deleting record for 'bond_james' with value  [Shaken, not stirred Martinis Women]
// Iterating values of map using range:
// Record for moneypenny_miss
//         0 ==> James Bond
//         1 ==> Literature
//         2 ==> Computer Science
// Record for no_dr
//         0 ==> Being evil
//         1 ==> ice cream
//         2 ==> Sunsets
// ===========================

func question_10() {
	fmt.Printf("\n\n============= Q-10 ==============\n")
	mp := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "ice cream", "Sunsets"},
	}
	if v, ok := mp["bond_james"]; ok {
		fmt.Println("deleting record for 'bond_james' with value ", v)
		delete(mp, "bond_james")
	}

	fmt.Printf("Iterating values of map using range: \n")
	for k, v := range mp {
		fmt.Printf("Record for %v\n", k)
		for i, e := range v {
			fmt.Printf("\t%d ==> %v\n", i, e)
		}
	}
	fmt.Printf("===========================\n")
}

func slice_1() {
	// multi-dimensional slice
	jb := []string{"James", "Bond", "choco", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "money", "straw", "nuts"}
	fmt.Println(mp)

	// multi-dimensional slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)

	/* o/p
	* [James Bond choco martini]
	* [Miss money straw nuts]
	* [[James Bond choco martini] [Miss money straw nuts]]
	 */
}

func map_1() {
	//m := map[string][int] {}
	fmt.Printf("\n\n============= MAPS ==============\n")
	// this below is a composite literal
	m := map[string]int{
		"James":           32,
		"Miss moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"]) // 32

	// if you enter the key of whose value is not stores in the map for that key
	// then it will return 0 value
	fmt.Println(m["hello"]) // 0

	// to chck that
	v, ok := m["hello"]
	fmt.Println(v, ok) // 0, false

	// for above we should write like this in go idiomitic way
	if vv, okk := m["hello"]; !okk {
		// o/p - hello does not exist. And this shoud be 0 0
		fmt.Printf("hello does not exist. And this shoud be 0 %v\n", vv)
	}

	// add an entry into map
	m["todd"] = 33
	fmt.Println(m) //map[James:32 Miss moneypenny:27 todd:33]

	/* o/p
	* James 32
	* Miss moneypenny 27
	* todd 33
	 */
	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete an entry from the map
	delete(m, "James")
	delete(m, "we can delete key even when it does not exist")
	fmt.Println(m)

	// so we should do it like this.
	if _, ok := m["James"]; ok {
		delete(m, "James")
	}

	fmt.Printf("===========================\n\n")
}

func main() {
	// array declaration
	// but pretty much use slices itself as per the golang reference too
	// the length is part of the array type
	var x [5]int
	var y [6]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Printf("%T %T\n", x, y)

	// SLICES
	//xx := type{values} // composite leteral
	// composite literal allows to do combine the data of the same type
	xx := []int{4, 5, 7, 8, 42}
	fmt.Println(xx, xx[0], len(xx), cap(xx))

	for i, v := range xx {
		fmt.Println("range over slice ", i, v)
	}

	for i := 0; i < len(xx); i++ {
		fmt.Println("using simple indexing for loop ", i, xx[i])
	}

	xx = xx[1:]
	fmt.Println(xx[1:])

	xx = append(xx, 77, 88, 99, 100)
	fmt.Println(xx)

	yy := []int{234, 456, 989}
	// take whole bunch of values from slice yy and put it all right here.
	// this is appending slice to a slice.
	// Note that alone yy does not work
	// note that append has the second argument as a variadic paramter
	// which is ...[]T (if ... is before the type then it means it can take
	// any number of these params.
	xx = append(xx, yy...)
	fmt.Println(xx)

	//delete from a slice. This is the idiomitic go way of doing it
	xx = append(xx[:2], xx[4:]...)
	fmt.Println(xx)

	xxm := make([]int, 10, 12)
	xxm[0] = 42
	xxm[9] = 442
	fmt.Println(xxm, len(xxm), cap(xxm))

	/* Output of this below. Hence capacity is imp. to avoid the copying
	* of array to a different array of double capacity.
	* [42 0 0 0 0 0 0 0 0 442] 10 12
	* [42 0 0 0 0 0 0 0 0 442 343] 11 12
	* [42 0 0 0 0 0 0 0 0 442 343 344] 12 12
	* [42 0 0 0 0 0 0 0 0 442 343 344 345] 13 24
	 */

	xxm = append(xxm, 343)
	fmt.Println(xxm, len(xxm), cap(xxm))
	xxm = append(xxm, 344)
	fmt.Println(xxm, len(xxm), cap(xxm))
	xxm = append(xxm, 345)
	fmt.Println(xxm, len(xxm), cap(xxm))

	slice_1()
	map_1()

	question_1()
	question_2()
	question_3()
	question_4()
	question_5()
	question_6()
	question_7()
	question_8_9()
	question_10()
}
