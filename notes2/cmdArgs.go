package notes2

import (
	"flag"
	"fmt"
	"os"
)

func CmdArgs() {
	// 接收的全部flag
	fmt.Printf("接收到了%d个参数\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("第%v个参数：%v\n", i, v)
	}

	// 接收有指定flag的参数
	vPr := flag.Bool("v", false, "版本") //-v 默认false|true
	var str string
	flag.StringVar(&str, "u", "", "用户名")      //-u 存到变量中
	flag.Func("f", "", func(s string) error { //解析到 -f 就会执行里面的函数
		fmt.Println("f=", s)
		return nil
	})

	flag.Parse()
	if *vPr {
		fmt.Println("当前版本是V0.0.0")
	}
	fmt.Printf("用户名%v\n", str)

}
