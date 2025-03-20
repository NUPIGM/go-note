package notes

import "fmt"

func Point() {

	var str string = "字符串"
	str2 := &str
	*str2 = "新字符串"
	str3 := str
	str3 = "字符串3"
	fmt.Println(&str, str)
	fmt.Println(str2, *str2)
	fmt.Println(&str3, str3)
	// 普通指针
	a := "hello"
	b := &a
	*b = "world"
	// fmt.Println(a)
	// 数组指针
	arr := [5]int{1, 2, 3, 4, 5}
	arrp := &arr
	arrp[3] = 0

	// fmt.Println(arr)
	var str4 = "hello"
	hello(&str4)

}

// 指针传参
func hello(str *string) {
	fmt.Println(*str)
}
