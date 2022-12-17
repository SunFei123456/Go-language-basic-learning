// Go 语言函数值传递值
/*传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。*/
package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)

	/* 通过调用函数来交换值 */
	swap1(a, b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

/* 定义相互交换值的函数 */
func swap1(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值  temp = 100*/
	x = y    /* 将 y 值赋给 x  x = 200*/
	y = temp /* 将 temp 值赋给 y y = 100*/

	return temp
}
