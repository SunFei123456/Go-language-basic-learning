package main

/*`1.声明数组
`	定义了数组 balance 长度为 10 类型为 float32：
`	var balance [10] float32  */

/* 2. 声明并定义
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}*/

/* 3.如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} */

/* 4.如果设置了数组的长度，我们还可以通过指定下标来初始化元素：
     将索引为 1 和 3 的元素初始化
     balance := [5]float32{1:2.0,3:7.0}

   5.初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
     如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
     balance[4] = 50.0*/
