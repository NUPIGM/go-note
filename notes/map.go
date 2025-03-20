package notes

import "fmt"

// map数据类型
func Map() {
	// 两种创建方法
	a := map[int]string{}
	a[1] = "a"
	a[2] = "b"
	b := make(map[string]interface{})
	b["a"] = "a"
	fmt.Println(a)
	fmt.Println(b)
	// 不同数据类型赋值
	c := map[string]interface{}{}
	c["a"] = "a"
	c["b"] = 2
	c["c"] = true
	// 循环出结果
	for k, v := range c {
		fmt.Println(k, v)
	}
}
