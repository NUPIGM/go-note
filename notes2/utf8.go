package notes2

import (
	"fmt"
	"unicode/utf8"
)

func Utf8Count() {
	var str int = utf8.RuneCountInString("hello,世界")
	fmt.Println(str)
}

func Utf8Valid() {
	str := "hello,世界"
	boolen := utf8.ValidString(str[:len(str)-1])
	fmt.Println(boolen)

}
