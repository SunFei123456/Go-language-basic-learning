package main

// import "fmt" 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
import "fmt"

func main() {
	const name1 string = "李四" // 显示定义

	const name = "张三" // 隐式定于

	const a, b, c = "张三", 18, true

	fmt.Println(name1)
	fmt.Println(name)
	fmt.Println(a, b, c)

}
