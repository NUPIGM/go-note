package main

import (
	"log"
)

func test() {
	// 数组
	a := [5]int{1, 2, 3, 4}
	// 切片
	b := a[2:]
	c := append(b, 2)
	// 不同方法
	d := new([1]int) // 返回一个指针
	e := make([]string, 6)
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Println(&d)
	log.Println(e)

}
