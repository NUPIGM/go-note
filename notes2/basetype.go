package notes2

import (
	"fmt"
	"strconv"
)

func Typeto() {
	var str string = "999"
	i, _ := strconv.Atoi(str)
	fmt.Printf("Atoi函数，字符串 %T -> 整型 %T\n", str, i)

	var integer int = 999
	a := strconv.Itoa(integer)
	fmt.Printf("Itoa函数，整型 %T -> 字符串 %T\n", integer, a)

	var str2 string = "123"
	b, _ := strconv.ParseBool(str2)
	fmt.Printf("ParseBool函数，字符串 %T -> 布尔 %T\n", str, b)

	var bol bool = true
	a2 := strconv.FormatBool(bol)
	fmt.Printf("FormatBool函数，布尔 %T -> 字符串 %T\n", bol, a2)

}
