package main

import "fmt"

func main() {
	str := "hello world"
	fmt.Println(len(str))
	// 循环遍历字符串每个 字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
	// 	for range 循环遍历数组,切片,,,,,,
	// 返回下标和数值
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c", v)

	}
}
