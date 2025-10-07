package notes

import (
	"fmt"
)

func Slice() {
	// 数组
	var a [5]int = [5]int{1, 2, 3, 4}
	fmt.Println("a", a)
	// 切片 本身不储存数据，只引用数组
	b := a[2:]
	b[2] = 5
	c := append(b, 6)
	// 不同方法
	d := new([1]int) // 返回一个指针
	e := make([]string, 6)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("&d", &d)
	fmt.Println("e", e)
	fmt.Println("a", a)

}
