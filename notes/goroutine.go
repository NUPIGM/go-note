package notes

import (
	"fmt"
	"sync"
	"time"
)

func Goroutine() {
	// 进程等待
	var wg sync.WaitGroup
	go myfun(&wg)

	for i := range 10 {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 500)
	}
	// 等待期间不会执行后面的代码，
	// 线程没Done ，Wait会报错
	wg.Wait()
	fmt.Println("最后执行")

}

func myfun(wg *sync.WaitGroup) {
	// 开启一个线程
	wg.Add(1)
	time.Sleep(time.Second)
	fmt.Println("Go 多协程run")
	// 完成该线程
	wg.Done()
}
