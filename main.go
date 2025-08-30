package main

import "fmt"

func main() {
	// 数组：长度是类型的一部分
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr, len(arr), cap(arr))
	arr = append(arr, 4)
	fmt.Println(arr, len(arr), cap(arr))
	//invalid append: argument must be a slice; have arr (variable of type [3]int)
	// 切片：没有固定长度
	//2. 存储方式不同

	// 数组：一块连续的内存，值类型，大小固定。

	// 切片：是对数组的一层抽象，内部结构包含：

	// 指向底层数组的指针

	// 长度（len）

	// 容量（cap）

	// 所以 切片本质上是对数组的引用。
	var s []int = []int{1, 2, 3}
	fmt.Println(s, len(s), cap(s))
}
