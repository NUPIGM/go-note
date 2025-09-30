package notes

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 管道配合协程一起使用
// 计算素数
func Prime() {
	var c int
	var c1 chan int = make(chan int, 100)

	for i := 2; i < 100001; i++ {
		go func() {
			for j := 2; j < i; j++ {
				if i%j == 0 {
					return
				}
			}
			c1 <- i
		}()
	}
done:
	for {
		select {
		case v := <-c1:
			c++
			fmt.Printf("%v\t", v)
		default:
			close(c1)
			break done
		}
	}
	fmt.Println("总数", c)
}

func Chan1() {
	// chanel
	c1 := make(chan int, 5)
	go func() {
		for i := range 10 {
			c1 <- i
		}
	}()
	for range 10 {
		log.Println(<-c1)
	}
}

func Chan2() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 1
	ch3 <- 1
done:
	for {
		select {
		case <-ch1:
			log.Println("c1")
		case <-ch2:
			log.Println("c2")
		case <-ch3:
			log.Println("c3")
		default:
			log.Println("none")
			break done
		}
	}

}

// 可读，可写的chanel
func Chan3() {
	c := make(chan int, 1)
	var readc <-chan int = c
	var writec chan<- int = c
	var wg sync.WaitGroup
	wg.Add(2)
	go set(writec, &wg)
	go get(readc, &wg)
	wg.Wait()
}

func set(writec chan<- int, wg *sync.WaitGroup) {
	for i := range 10 {
		writec <- i
	}
	close(writec)
	wg.Done()
}
func get(readc <-chan int, wg *sync.WaitGroup) {
	for range 10 {
		log.Println(<-readc)
	}
	wg.Done()
}

func TimeChan() {
	var intchan chan int = make(chan int)
	select {
	case <-intchan:
		fmt.Println("收到验证码")
	case <-time.After(time.Second * 2):
		fmt.Println("验证码以过期")
	}
}
