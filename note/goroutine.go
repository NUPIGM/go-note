package main

import (
	"log"
	"sync"
)

func main() {
	// 进程等待
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go myfun(&wg)
	// wg.Wait()

}

func myfun(wg *sync.WaitGroup) {
	log.Println("run")
	wg.Done()
}
