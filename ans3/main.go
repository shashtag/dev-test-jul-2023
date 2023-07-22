package main

import "fmt"

func main() {

	// driver code for iterative version  - change as you like to test out the code.
	itr(10.32323, 20.234234)
	itr(320.823, 230.93)
	itr(0.1, 230.93)
	itr(1, 29.50)
	itr(2, 330.61)
	itr(43543)
	itr(44)
}

func itr(n float64, args ...float64) {

	var cur float64

	// handle panic/error - needed for GoLang (not part of the code in question)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s : n = %f \n", r, n)
		}
	}()

	// base case 1
	if len(args) == 0 {
		cur = 0
	} else {
		cur = args[0]
	}
	// base case 2
	if n < 2 {
		panic("Invalid input")
	}

	// iterative version of recursive case
	for ; n > 2; n-- {
		cur = cur + 1/(n*(n-1))
	}

	// print result
	fmt.Println(1/n + cur)
}
