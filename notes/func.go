package notes

import (
	"fmt"
)

func Func() {
	// [TODO]

	ret1, ret2 := a(1, "a")
	fmt.Println(ret1, ret2)

	bb := []string{"1", "2", "3"}
	b(999, bb...)
	// 延时调用 defer,函数结束才会运行defer
	defer c()(3)
	fmt.Println("在3后面")

	// 自执行函数
	func() {
		fmt.Println("自执行函数")
	}()

}

// 闭包函数
func c() func(int) {
	return func(num int) {
		fmt.Println(num)
	}
}

// 不定向参数
func b(data1 int, data2 ...string) {
	fmt.Println(data1, data2)
	for k, v := range data2 {
		fmt.Println(k, v)
	}
}

// 常用函数声明 -> 函数名（入参）（出参）
func a(data1 int, data2 string) (ret1 int, ret2 string) {
	ret1 = data1
	ret2 = data2
	return
}
