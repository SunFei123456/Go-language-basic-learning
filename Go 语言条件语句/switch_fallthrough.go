package main

import "fmt"

func main() {
	var score = 90
	switch score {
	case 90:
		fmt.Println("评级为:A")
		fallthrough // case穿透的,不管下一个条件满不满足,直接执行 *很少用
	case 80:
		if score == 90 {
			break // 终止case穿透
		}
		fmt.Println("评级为:B")
	case 70, 60:
		fmt.Println("评级为:C")
	default:
		fmt.Println("不及格,继续加油")

	}

}
