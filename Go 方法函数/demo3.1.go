package main

import "fmt"

func main() {
	// 	切片 -- 可扩容的数组
	s1 := []int{1, 2, 3, 4}
	updata(s1)
	// 	通过方法修改了参数的值的变化也影响到了实参的变化  因为切片是引用类型
	fmt.Println("通过方法修改之后的切片的值为:", s1)

}
func updata(s2 []int) {
	fmt.Println("传递前的切片为:", s2)
	s2[0] = 100
	fmt.Println("修改之后的切片为:", s2)

}
