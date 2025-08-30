package main

import "fmt"

func f() int {
	return 0
}

func g(n int) int {
	return n
}

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	fmt.Println(x, y) // compile error: x and y are not visible here
}
