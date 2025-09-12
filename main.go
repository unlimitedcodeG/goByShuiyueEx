package main

import "fmt"

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	// True if equal is written incorrectly.
	equal(map[string]int{"A": 0}, map[string]int{"B": 42})

	fmt.Println(ages)
	fmt.Println(ages["alice"])
	fmt.Println(ages["charlie"])
	fmt.Println(ages["bob"])
	delete(ages, "alice") // remove element ages["alice"]
	fmt.Println(ages)
	ages["bob"] = ages["bob"] + 1 // happy birthday!
	fmt.Println(ages["bob"])
}
