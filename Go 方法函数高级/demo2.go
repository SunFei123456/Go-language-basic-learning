// 匿名函数
package main

import "fmt"

func main() {
	f3()
	f4 := f3 // 不叫括号是地址是变量 加了括号是方法 是函数
	f4()

	// 匿名函数
	f5 := func() {
		fmt.Println("我是f5匿名函数")
	}
	f5()
	// 匿名函数自己调用自己
	f6 := func(a, b int) int {
		fmt.Println("我是匿名函数")
		return a * b

	}(2, 2)
	fmt.Println("a×b的值为:", f6)

}
func f3() {
	fmt.Println("你好")
}
