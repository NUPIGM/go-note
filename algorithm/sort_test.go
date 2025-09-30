package algorithm_test

import (
	"fmt"
	"main/algorithm"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func randArray(n int) []int {
	var array []int
	var wg sync.WaitGroup
	for i := range n + 1 {
		wg.Add(1)
		go func() {
			array = append(array, i)
			wg.Done()
		}()
	}
	wg.Wait()
	return array
}
func randNum(n int) []int {
	var arr []int = make([]int, n)
	seedNum := time.Now().UnixNano()
	for i := range n {
		rand.Seed(seedNum)
		arr[i] = rand.Intn(n + 1)
		seedNum++
	}
	return arr
}
func TestBubbleSort(t *testing.T) {
	num := randArray(100)
	fmt.Println(num)
	algorithm.SelectionSort(num)
	fmt.Println(num)
}
