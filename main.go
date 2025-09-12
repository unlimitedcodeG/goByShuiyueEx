package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// 提示用户输入
	fmt.Println("请输入文本（输入完成后按 Ctrl+D 结束）:")

	// 创建一个map来存储单词频率
	wordFreq := make(map[string]int)

	// 创建scanner从标准输入读取
	input := bufio.NewScanner(os.Stdin)

	// 设置按单词分割而不是按行分割
	input.Split(bufio.ScanWords)

	// 扫描每个单词并统计频率
	for input.Scan() {
		word := input.Text()
		wordFreq[word]++
	}

	// 检查扫描过程中是否有错误
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	// 为了有序输出，将单词提取到切片中并排序
	var words []string
	for word := range wordFreq {
		words = append(words, word)
	}
	sort.Strings(words)

	// 输出每个单词及其频率
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, wordFreq[word])
	}
}
