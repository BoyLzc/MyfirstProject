package main

import (
	"fmt"
	"strconv"
)

func isPalindromeNumber(num int) bool {
	// 将整数转换为字符串
	str := strconv.Itoa(num)
	// 检查字符串是否是回文
	// 头尾开始对齐
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	num := 12321
	if isPalindromeNumber(num) {
		fmt.Printf("%d 是回文数\n", num)
	} else {
		fmt.Printf("%d 不是回文数\n", num)
	}
}
