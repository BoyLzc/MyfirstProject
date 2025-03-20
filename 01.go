package main

import "fmt"

func main() {
	// 定义变量
	name := "小明"   // 姓名，string 类型
	age := 23      // 年龄，int 类型
	gender := true // 性别，true 表示男，false 表示女，bool 类型

	// 打印信息
	fmt.Println("姓名:", name)
	fmt.Println("年龄:", age)
	if gender {
		fmt.Println("性别: 男")
	} else {
		fmt.Println("性别: 女")
	}
}
