package notes

import "fmt"

func Array() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)
	// ...自适应数组长度
	var b = [...]int{1, 2, 3, 4}
	fmt.Println(b)
	//遍历循环
	for i, v := range a {
		fmt.Println(i, v)
	}
	//多维数据
	var c [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, v := range c {
		for _, t := range v {
			fmt.Print(t)
		}
		fmt.Println()
	}
}
