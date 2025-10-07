package algorithm

import (
	"math/rand"
	"sync"
	"time"
)

// 打乱的有序数组
func RandArray(n int) []int {
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
	array = WishPock(array)
	return array
}

// 洗牌/打乱
func WishPock(arr []int) []int {
	rand.Seed(time.Now().UnixNano())
	// 打乱数组
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	return arr
}

// 打乱的无序数组
func RandNum(n int) []int {
	var arr []int = make([]int, n)
	seedNum := time.Now().UnixNano()
	for i := range n {
		rand.Seed(seedNum)
		arr[i] = rand.Intn(n + 1)
		seedNum++
	}
	return arr
}
