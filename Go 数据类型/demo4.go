// 数据类型转换
package main

import "fmt"

func main() {
	a := 3    // int类型
	b := 5.00 // float类型
	// 	将float类型的数据转换成int类型的
	c := int(b)
	fmt.Printf("a的类型为:%T,值为:%d\n", a, a)
	fmt.Printf("c的类型为:%T,值为:%d\n", c, c)
	// 	将int类型转换为float类型的数值
	d := float64(a)
	fmt.Printf("d的类型为:%T,值为:%f\n", d, d)

}
