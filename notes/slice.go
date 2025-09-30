package notes

import (
	"fmt"
)

func Slice() {
	// 数组
	var a [5]int = [5]int{1, 2, 3, 4}
	fmt.Println(a)
	// 切片 本身不储存数据，只引用数组
	b := a[2:]
	b[2] = 5
	c := append(b, 2)
	// 不同方法
	d := new([1]int) // 返回一个指针
	e := make([]string, 6)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(&d)
	fmt.Println(e)
	fmt.Println(a)

}
