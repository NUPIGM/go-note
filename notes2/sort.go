package notes2

import (
	"fmt"
	"sort"
)

// 整数排序
func IntSort() {
	s := []int{2, 4, 7, 1, 5}
	sort.Ints(s)
	fmt.Println(s)
}

// 字母排序
func StrSort() {
	s := []string{"hello", "world", "good", "morning"}
	sort.Strings(s)
	fmt.Println(s)
}

// 排序搜索
func SearchIntSort() {
	s := []int{2, 4, 6, 8, 9}
	p1 := sort.SearchInts(s, 8) //有时，返回已有的下标
	p2 := sort.SearchInts(s, 5) //没有时，返回适合位置的下标前
	fmt.Println(p1)
	fmt.Println(p2)
}

type peoson struct {
	name string
	age  int
}

// 自定义排序
func SliceSort() {

	p := []peoson{{"john", 18}, {"siri", 8}, {"mick", 23}}
	sort.Slice(p, func(i, j int) bool {
		return p[i].age > p[j].age
	})
	fmt.Println(p)
}

// 自定义查找
func SearchSort() {
	s := []int{2, 4, 6, 8, 9}

	p := sort.Search(len(s), func(i int) bool {
		return s[i] > 6
	})
	fmt.Println(p)
}
