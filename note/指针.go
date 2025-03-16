package main

import "fmt"

func main() {
	// 普通指针
	a := "hello"
	b := &a
	*b = "world"
	// fmt.Println(a)
	// 数组指针
	arr := [5]int{1, 2, 3, 4, 5}
	arrp := &arr
	arrp[3] = 0

	// fmt.Println(arr)
	var str = "hello"
	hello(&str)

}

// 指针传参
func hello(str *string) {
	fmt.Println(*str)
}
