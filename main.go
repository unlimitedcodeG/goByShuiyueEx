package main

import "fmt"

// 方法1: 切片重组 (保持顺序)
func removeBySlicing(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// 方法2: 使用copy函数 (保持顺序)
func removeByCopy(s []int, index int) []int {
	copy(s[index:], s[index+1:])
	return s[:len(s)-1]
}

// 方法3: 交换删除 (不保持顺序，但效率高)
func removeBySwap(s []int, index int) []int {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

// 方法4: 删除多个连续元素
func removeRange(s []int, start, end int) []int {
	return append(s[:start], s[end:]...)
}

// 方法5: 按条件删除元素
func removeByCondition(s []int, condition func(int) bool) []int {
	result := s[:0] // 重用底层数组
	for _, v := range s {
		if !condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// 方法6: 删除所有指定值
func removeAllValue(s []int, value int) []int {
	result := s[:0]
	for _, v := range s {
		if v != value {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("=== Go中删除slice元素的多种方法 ===")

	// 原始数据
	original := []int{1, 2, 3, 4, 5}

	// 方法1: 切片重组
	s1 := make([]int, len(original))
	copy(s1, original)
	s1 = removeBySlicing(s1, 2) // 删除索引2的元素(值为3)
	fmt.Printf("方法1 - 切片重组: %v\n", s1)

	// 方法2: 使用copy
	s2 := make([]int, len(original))
	copy(s2, original)
	s2 = removeByCopy(s2, 2)
	fmt.Printf("方法2 - 使用copy: %v\n", s2)

	// 方法3: 交换删除
	s3 := make([]int, len(original))
	copy(s3, original)
	s3 = removeBySwap(s3, 2)
	fmt.Printf("方法3 - 交换删除: %v (顺序可能改变)\n", s3)

	// 方法4: 删除范围
	s4 := make([]int, len(original))
	copy(s4, original)
	s4 = removeRange(s4, 1, 3) // 删除索引1到2的元素
	fmt.Printf("方法4 - 删除范围[1:3): %v\n", s4)

	// 方法5: 按条件删除
	s5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s5 = removeByCondition(s5, func(n int) bool { return n%2 == 0 }) // 删除偶数
	fmt.Printf("方法5 - 删除偶数: %v\n", s5)

	// 方法6: 删除所有指定值
	s6 := []int{1, 2, 3, 2, 4, 2, 5}
	s6 = removeAllValue(s6, 2) // 删除所有值为2的元素
	fmt.Printf("方法6 - 删除所有2: %v\n", s6)

	fmt.Println("\n=== 性能对比说明 ===")
	fmt.Println("• 切片重组: 保持顺序，需要移动元素，O(n)时间复杂度")
	fmt.Println("• copy函数: 保持顺序，需要移动元素，O(n)时间复杂度，可能更高效")
	fmt.Println("• 交换删除: 不保持顺序，O(1)时间复杂度，最高效")
	fmt.Println("• 条件删除: 适合批量删除，只遍历一次")
}
