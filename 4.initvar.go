package main

import (
	"fmt"
)

func main() {
	// 1.字符串
	var a = "Hello"
	fmt.Println("a =", a)

	// 2. 整型
	var b int
	// int初始化值为0
	fmt.Println("b =", b)
	b = 1
	fmt.Println("b =", b)

	// 3. 布尔型
	var c bool
	// bool初始化值为false
	fmt.Println("c =", c)
	c = true
	fmt.Println("c =", c)

	// 批量声明变量
	var (
		d string
		e int
		f bool
		g float32
		h float64
	)

	// 批量赋值
	d, e, f, g, h = "Hello", 233, true, 3.14, 3.1415926

	println(d, e, f, g, h)

}
