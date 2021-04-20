package main

import (
	"fmt"
)

func main() {
	// 普通常量
	const c1 = "HelloWOlrd"
	fmt.Println("c1 :", c1)
	// 枚举定义常量
	const (
		Unknow = 0
		Female = 1
		Male   = 2
	)

	fmt.Println(Male)
}
