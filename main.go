package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	select {
	case ch <- 1:
		fmt.Println("send success")
	default:
		fmt.Println("send failed")
	}
}
