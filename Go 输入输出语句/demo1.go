package main

import "fmt"

func main() {
	// 	输出语句
	// fmt.Println() // 打印并换行
	// fmt.Printf()  // 格式化输出
	// fmt.Print()   // 打印输出

	// 	输入语句 -- 用户
	// fmt.Scanln() // 接收输入并换行
	// fmt.Scanf()  // 接收输入 格式化输出
	// fmt.Scan()   // 接收输入

	var name = ""
	var age = "0"
	var school = ""
	fmt.Println("请输出你的姓名")
	fmt.Scanln(&name)
	fmt.Println("请输出你的年龄")
	fmt.Scanln(&age)
	fmt.Println("请输出你的学校")
	fmt.Scanln(&school)
	fmt.Printf("姓名为:%s\n"+"年龄为:%s\n"+"学校为:%s\n", name, age, school)
}
