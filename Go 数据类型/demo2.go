// 整型浮点类型
package main

import "fmt"

func main() {
	// 定义一个整形
	// byte uint8
	// rune int32

	// 	int int64
	var age int = 18
	fmt.Printf("age的数据类型为:%T,值为:%d\n", age, age)

	// 	定义一个浮点型
	var money float64 = 3.19
	fmt.Printf("money的数据类型为:%T,值为:%f\n", money, money) // float64 默认小数点为6位小数
	// 	通过在格式化打印的时候在值格式化添加指定格式可以控制float类型的数值小数点后保留几位
	fmt.Printf("money的数据类型为:%T,值为:%.1f\n", money, money) // float64 默认小数点为6位小数

	// 	float64尽量使用	float64 来定义浮点类型的数据 防止数据丢失
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1 =", num1, "num2 =", num2)

}
