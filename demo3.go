package main

import "fmt"

func main() {
	//1.定义变量类型
	var (
		name    string
		age     int
		address string
	)
	//2.给变量进行赋值
	name = "张三"
	age = 18
	address = "上海"
	// string : 默认值为 空
	// int : 默认值为  0
	fmt.Println(name, age, address)
}
