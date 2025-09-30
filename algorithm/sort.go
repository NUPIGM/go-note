package algorithm

// 冒泡排序
func BubbleSort(n []int) {
	lenIndex := len(n) - 1
	for i := 0; i < lenIndex; i++ {
		for j := 0; j < lenIndex-i; j++ {
			if n[j] > n[j+1] {
				a := n[j]
				n[j] = n[j+1]
				n[j+1] = a
			}
		}
	}
}

// 选择排序
func SelectionSort(n []int) {
	lenIndex := len(n) - 1
	for i := 0; i < lenIndex; i++ {
		max := lenIndex - i
		for j := 0; j < i; j++ {
			if n[j] < n[max] {
				a := n[j]
				n[j] = n[max]
				n[max] = a
			}
		}
	}
}
