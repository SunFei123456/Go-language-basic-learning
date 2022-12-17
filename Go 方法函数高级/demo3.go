// 回调函数
package main

import "fmt"

func main() {
	r1 := add(1, 2)
	fmt.Println(r1)

	fmt.Println("调用回调函数得到值", oper(1, 2, add))

}

// 一个函数作为另外一个函数的参数
func oper(a, b int, fun func(int, int) int) int {
	r := fun(a, b)
	return r
}
func add(a, b int) int {
	return a + b
}
