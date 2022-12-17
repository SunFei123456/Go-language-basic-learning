// 字符串类型
package main

import "fmt"

func main() {
	// "":为字符串
	str := "张三"
	// '' : 为字符  -- 整型,-ASCLL码表
	// 扩展:
	// 所有中国字的编码表 : GBK
	// 世界统一的编码表为:Unicode 编码表
	var s = 'A'
	fmt.Printf("str的数据类型为:%T,值为:%s\n", str, str)
	fmt.Printf("s的数据类型为:%T,值为:%d\n", s, s)

	fmt.Println("=====================================")
	// 	字符串的拼接 +
	fmt.Println("hello" + "你好")
	fmt.Println("=====================================")

	// 	转义字符
	// \n 换行输出
	fmt.Println("hello\n你好")
	// 	\转义字符
	fmt.Println("hello\"你好")
	// \t 空格
	fmt.Println("hello\t你好")

}
