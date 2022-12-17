package main

import "fmt"

func main() {
	//我们只需要获取一个变量的值的时候,第二个变量可以定义为匿名变量_
	a, _ := test()
	fmt.Println(a)
}
func test() (int, int) {
	return 100, 200
}
