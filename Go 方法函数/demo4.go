// 可变参数
package main

import "fmt"

func main() {

	fmt.Println(addnums(1, 2, 3, 4, 5))
}

// 可变参数 当参数不知道到底有多少个 可以使用... 数据类型来表示
func addnums(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum

}
