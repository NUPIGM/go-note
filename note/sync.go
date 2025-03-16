package main

import (
	"fmt"
	"sync"
)

func main() {
	//只能运行一次
	o := &sync.Once{}

	o.Do(func() {
		fmt.Println("只执行一次")
	})
	o.Do(func() {
		fmt.Println("hello")
	})

	//map类型
	m := &sync.Map{}
	m.Store(1, 1)
	m.Store(2, 2)
	m.Store(3, 3)
	m.Load(1) //不知道干嘛用
	m.Delete(2)
	//顺序随机
	m.Range(func(key, value any) bool {
		fmt.Println("map:", key, value)
		return true
	})

	//pool
	p := &sync.Pool{}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)
	p.Put(5)
	for i := 0; i < 5; i++ {
		fmt.Println("pool池:", p.Get())
	}
}
