package notes2

import (
	"errors"
	"fmt"
)

func Err() {
	e := errors.New("错误")
	fmt.Println(e)
}
