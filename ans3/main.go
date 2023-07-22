package main

import "fmt"

func main() {

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
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s : n = %f \n", r, n)
		}
	}()

	if len(args) == 0 {
		cur = 0
	} else {
		cur = args[0]
	}

	if n < 2 {
		panic("Invalid input")
	}

	for ; n > 2; n-- {
		cur = cur + 1/(n*(n-1))
	}

	fmt.Println(1/n + cur)
}
