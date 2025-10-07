package algorithm

// 冒泡排序
func BubbleSort(n []int) {
	lenIndex := len(n) - 1
	for i := 0; i < lenIndex; i++ {
		for j := 0; j < lenIndex-i; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
}

// 选择排序
func SelectionSort(n []int) {
	lenIndex := len(n) - 1
	for i := 0; i < len(n); i++ {
		max := lenIndex - i
		for j := 0; j < max; j++ {
			if n[j] > n[max] {
				n[j], n[max] = n[max], n[j]
			}
		}
	}
}

// 插入排序

//二分查找
