package main

import "fmt"

func main() {
	fmt.Println(max(20, 100))
	fmt.Println(swap("sunfei", "zhangsan"))
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*函数交换并且返回多个结果值*/
func swap(x, y string) (string, string) {
	return y, x
}
