package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"runtime"
	"sync"
)

// Concerruncy chapter
var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("GORoutine:\t", runtime.NumGoroutine())

	// without wait group foo could never run, since main exits.
	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
	//race_condition_example()
	race_condition_example_mutex()
	concurrency_golangbootcamp()
	example_tree_walk()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func race_condition_example() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	cnt := 0 // scope of this variable is only until this function and until all the go routines are finished

	const gs = 100
	var wg_1 sync.WaitGroup
	wg_1.Add(gs)

	for i := 0; i < gs; i++ {
		// this happens because multiple go routines are accessing
		// meaning reading modifying and writing the same variable
		// cnt in parallel.
		// This is same as RMW (read-modify-write cycle), which needs
		// to be protected.
		go func() {
			v := cnt
			// to better replicate the race you can uncomment this
			// runtime.Gosched()
			v++
			cnt = v
			wg_1.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg_1.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", cnt)

	// final answer of count varies sometimes := 99, 95, 100, 98 etc...
	// go run -race main.go
	// Count: 99
	// Found 2 data race(s)
	// exit status 66
}

func race_condition_example_mutex() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	cnt := 0 // scope of this variable is only until this function and until all the go routines are finished

	const gs = 100
	var wg_1 sync.WaitGroup
	wg_1.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		// this happens because multiple go routines are accessing
		// meaning reading modifying and writing the same variable
		// cnt in parallel.
		// This is same as RMW (read-modify-write cycle), which needs
		// to be protected.
		// mutex examplg
		go func() {
			mu.Lock()
			v := cnt
			// to better replicate the race you can uncomment this
			// runtime.Gosched()
			v++
			cnt = v
			mu.Unlock()
			wg_1.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg_1.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Count:", cnt)

	// final answer of count varies sometimes := 99, 95, 100, 98 etc...
	// go run -race main.go
	// Count: 99
	// Found 2 data race(s)
	// exit status 66
	// go run -race main.go does not report any race with this code
}

func concurrency_golangbootcamp() {
	// http://www.golangbootcamp.com/book/concurrency
	//
	// e.g. of how to split the sum of array into two halves.

	c := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// send channel.
	// I remember as to channel
	// recieve or from channel should be declared as
	// ch <-chan int
	func_sum := func(ch chan<- int, vi ...int) {
		sum := 0
		for _, v := range vi {
			sum += v
		}
		ch <- sum
	}
	go func_sum(c, arr[:len(arr)/2]...)
	go func_sum(c, arr[len(arr)/2:]...)
	x, y := <-c, <-c
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("concurrent sum:", x+y, sum)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)
	// closing the channel so range can finish
	close(ch)
}

// recWalk walks recursively through the tree and push values to the channel
// at each recursion
func recWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		// send the left part of the tree to be iterated over first
		recWalk(t.Left, ch)
		// push the value to the channel
		ch <- t.Value
		// send the right part of the tree to be iterated over last
		recWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			// not the same size
			return false
		case !ok1:
			// both channels are empty
			return true
		case x1 != x2:
			// elements are different
			return false
		default:
			// keep iterating
		}
	}
}

func example_tree_walk() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

	// o/p is like below
	// 1       2       3       4       5       6       7       8       9       10
	// true
	// false

}
