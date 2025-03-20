package notes

import (
	"fmt"
	"sync"
	"time"
)

func Sync() {
	//只能运行一次
	o := &sync.Once{}

	o.Do(func() {
		fmt.Println("只执行一次")
	})
	o.Do(func() {
		fmt.Println("hello")
	})

	//map类型
	m := &sync.Map{}
	m.Store(1, "map1")
	m.Store(2, 2)
	m.Store(3, 3)
	m.Load(1) //获取值
	//无顺序
	m.Range(func(key, value any) bool {
		fmt.Println("map:", key, value)
		return true
	})

	// 并发池
	var pool = &sync.Pool{
		New: func() any {
			fmt.Println("新建对象")
			return new([]byte) // 这里创建的是 []byte 指针
		},
	}
	//pool，主要作用减少GCC内存回收
	// 从 Pool 获取对象
	obj1 := pool.Get().(*[]byte)
	fmt.Println("获取对象:", obj1)

	// 归还对象到 Pool
	// pool.get()没有指针，地址不一样
	pool.Put(obj1)

	// 再次获取对象（应该是之前归还的）
	obj2 := pool.Get().(*[]byte)
	fmt.Println("复用对象:", obj2)
}

// 锁
func Mux() {

	var wg sync.WaitGroup
	wg.Add(10)
	// 读锁
	// l := &sync.Mutex{}
	// 读写锁
	l := &sync.RWMutex{}
	go lockFun(l, &wg)
	go lockFun(l, &wg)
	go lockFun(l, &wg)
	go lockFun(l, &wg)
	go lockFun(l, &wg)
	go RLockFunc(l, &wg)
	go RLockFunc(l, &wg)
	go RLockFunc(l, &wg)
	go RLockFunc(l, &wg)
	go RLockFunc(l, &wg)
	wg.Wait()

}
func lockFun(lock *sync.RWMutex, wg *sync.WaitGroup) {
	// 阻塞
	lock.Lock()
	fmt.Println("attack")
	time.Sleep(time.Second)
	lock.Unlock()
	wg.Done()
}

func RLockFunc(lock *sync.RWMutex, wg *sync.WaitGroup) {
	//不阻塞
	lock.RLock()
	fmt.Println("heath")
	time.Sleep(time.Second)
	wg.Done()
	lock.RUnlock()
}
