package notes

import "fmt"

func util() func(n int) {
	i := 0
	return func(n int) {
		fmt.Println("接收", n)
		i++
		fmt.Println("被调用", i)
	}

}

// 先进后出
func Defer() {
	d := util()
	defer d(1)
	defer d(2)
	defer d(3)
}
