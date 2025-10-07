package algorithm_test

import (
	"fmt"
	"main/algorithm"
	"testing"
)

func TestWishPock(t *testing.T) {
	s := algorithm.WishPock([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(s)
}
func TestBubbleSort(t *testing.T) {
	num := algorithm.RandArray(35)
	fmt.Println(num)
	algorithm.SelectionSort(num)
	fmt.Println(num)
}
