package algorithm

import "fmt"

// 闭包
func B() func(n int) int {
	i := 0
	return func(n int) int {
		fmt.Println("接收了：", n)
		i++
		fmt.Println("被调用：", i)
		return i
	}
}
