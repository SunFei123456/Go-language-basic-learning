// 算数运算符
package main

import "fmt"

func main() {

	var a int = 20
	var b int = 10
	var c int
	// + 加
	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c) // 30
	// - 减
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c) // 10
	// * 乘以
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c) // 200
	// / 除以
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c) // 2
	// % 取余数
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c) // 0
	// ++ 自增
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a) // 21
	a = 21                            // 为了方便测试，a 这里重新赋值为 21
	// -- 自减
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a) // 20
}
