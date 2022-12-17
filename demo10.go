package main

import "fmt"

func main() {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println("---------------------------------------------")
	textA()
}

const (
	i = 1 << iota
	j = 3 << iota
	k
	l
)

// 公式 :  <<n==n(2^n)。
func textA() {
	fmt.Println("i=", i) //1*0
	fmt.Println("j=", j) // j*2^1
	fmt.Println("k=", k) //k*2^2
	fmt.Println("l=", l)
}
