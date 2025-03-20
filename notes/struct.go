package notes

import "fmt"

func Struct() {
	// 第一种申明方法
	var qm User
	fmt.Println(qm)
	qm.name = "qimiao"
	qm.age = 19
	qm.hobby = []string{"唱", "跳"}
	fmt.Println(qm)
	// 第二种申明方法,必需填写全部内容
	qm2 := User{"奇淼", 20, []string{"rap"}, Home{"where"}}
	fmt.Println(qm2)
	// 第三种申明方法
	qm3 := User{
		name:  "qm",
		age:   21,
		hobby: []string{"蓝球"},
	}
	fmt.Println(qm3)
	// 返回实例的地址
	qm4 := new(User)
	qm4.name = "qq"
	fmt.Println(qm4, "\n-------------")

	// struct作为参数
	userFunc(qm)

	// 指针修改内容
	var qm5 User
	qmq := &qm5
	qmq.name = "奇怪的奇淼"
	fmt.Println(qmq, "\n------------")

	// 绑定函数
	my := User{
		name: "我自己",
	}
	my.song("惊雷")
	fmt.Println("------------")

	// struct嵌套struct
	var qm6 User
	qm6.Home.addr = "北京"
	fmt.Println(qm6.Home.addr) //结果一样
	fmt.Println(qm6.addr)      //结果一样

}

type Home struct {
	addr string
}
type User struct {
	name  string
	age   int
	hobby []string
	Home
}

func (u User) song(name string) (ret string) {
	ret = "一计惊雷"
	fmt.Println(u.name + "，唱了一首" + name + "，真是" + ret)
	return
}

// struct作为参数
func userFunc(u User) {
	fmt.Println(u.name, "\n-----------")
}
