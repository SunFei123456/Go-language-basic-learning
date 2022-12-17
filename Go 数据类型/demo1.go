// 初始基本数据类型
package main

import "fmt"

func main() {
	// var 变量名 数据类型
	var isTrue bool
	println(isTrue)
	isTrue = true
	fmt.Printf("isTrue的数据类型为:%T,值为:%t", isTrue, isTrue)
}
