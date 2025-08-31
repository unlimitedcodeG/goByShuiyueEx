package main

import "fmt"

func f() int {
	return 0
}

func g(n int) int {
	return n
}

func main() {
	text := "hello"
	for _, ch := range text {
		upper := ch + 'A' - 'a'
		fmt.Printf("%c", upper)
	}
}
