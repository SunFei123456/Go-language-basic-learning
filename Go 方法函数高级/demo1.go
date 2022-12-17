// 函数本质的探讨
package main

import "fmt"

func main() {
	// 	f1 如果不加括号,就是一个变量
	fmt.Printf("%T\n", f1) // func() | func(int , int ) | func(int , int )int
	// 定义函数类型的变量
	var f2 func(int, int)
	f2 = f1         // 引用类型的 地址传递
	fmt.Println(f1) // 0xadfe20

	fmt.Println(f2) // 0xadfe20

	f2(1, 2)
}
func f1(a, b int) {
	fmt.Println(a, b)
}
