// 赋值运算符
package main

import "fmt"

func main() {
	var a int = 21
	var c int

	c = a //21
	fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c)
	// C = C + A
	c += a //21+21=42
	fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c)
	// C = C - A
	c -= a //42-21=21
	fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c)

	c *= a //21*21=441
	fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c)

	c /= a //441/21=21
	fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c)
	// 公式 :  a<<n==a(2^n)。
	c = 200

	c <<= 2 //200*2*2=800
	fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2 //200
	fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c)

	c &= 2
	fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c)
	//C = C ^ 2
	c ^= 2
	fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c)

	c |= 2
	fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c)

}
