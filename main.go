package main

import "fmt"

func main() {
	// 数组
	arr := [3]int{1, 2, 3}
	arrCopy := arr
	arrCopy[0] = 100
	fmt.Println(arr, arrCopy)

	//切片
	slice := []int{1, 2, 3}
	sliceCopy := slice
	sliceCopy[0] = 100
	fmt.Println(slice, sliceCopy)
}
