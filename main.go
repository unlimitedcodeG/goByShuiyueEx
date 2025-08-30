package main

import "fmt"

func main() {
	s := make([]int, 0, 10)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 2)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 3)
	fmt.Println(s, len(s), cap(s))
}
