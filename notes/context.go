package notes

import (
	"context"
	"fmt"
	"time"
)

// 设置超时（WithTimeout）
// 设定截止时间（WithDeadline）
// 手动取消（WithCancel）
// 传递请求范围的元数据（WithValue）
func Context() {
	// context上下文
	ctx, clear := context.WithCancel(context.Background()) //手动取消
	msg := make(chan int)
	go son(ctx, msg)
	for i := 0; i < 10; i++ {
		msg <- i
	}
	clear()
	time.Sleep(time.Second)
	fmt.Println("主进程结束了")
}

func son(ctx context.Context, msg chan int) {
	t := time.Tick(time.Millisecond * 500)
	for range t {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-ctx.Done():
			fmt.Println("son结束了")
			return
		}
	}
}
