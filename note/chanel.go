package main

import (
	"log"
	"sync"
)

func main() {
	chan3()

}

func chan1() {
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

func chan2() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 1
	ch3 <- 1
	select {
	case <-ch1:
		log.Println("c1")
	case <-ch2:
		log.Println("c1")
	case <-ch3:
		log.Println("c1")
	default:
		log.Println("none")
	}
}

// 可读，可写的chanel
func chan3() {
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
