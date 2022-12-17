package main

import "fmt"

//全局变量
var name string = "张三"

func main() {
	name = "李四"
	//局部变量
	var age int = 20
	fmt.Println(name, age)
	test1()
}
func test1() {
	name = "王五"
	fmt.Println(name)

}
