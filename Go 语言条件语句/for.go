package main

import "fmt"

func main() {
	// sum := 0
	/*	init： 一般为赋值表达式，给控制变量赋初值；
		condition： 关系表达式或逻辑表达式，循环控制条件；
		post： 一般为赋值表达式，给控制变量增量或减量。*/
	/*for num := 0; num < 10; num++ {
		sum = sum + num
	}*/
	// 这样写也可以，更像 While 语句形式
	/*	for sum <= 20 {
			sum = sum + sum
		}
		fmt.Println(sum)*/
	// 打印五五方阵
	for i := 0; i < 5; i++ {
		fmt.Println("* * * * *")
	}

}
