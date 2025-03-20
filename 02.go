package main

import (
	"fmt"
)

func countCharacters(s string) int {
	// 将字符串转换为 rune 数组
	runes := []rune(s)
	// 返回 rune 数组的长度
	return len(runes)
}

func main() {
	str1 := "Hello, 世界!"
	count := countCharacters(str1)
	fmt.Printf("字符串 \"%s\" 中的字符个数为: %d\n", str1, count)

	str2 := "世界!"
	count2 := countCharacters(str2)
	fmt.Printf("字符串 \"%s\" 中的字符个数为：: %d\n", str2, count2)
}
