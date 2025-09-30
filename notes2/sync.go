package notes2

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

func Waitting() {
	var wg sync.WaitGroup
	for i := 2; i < 100001; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 2; j < i; j++ {
				if i%j == 0 {
					return
				}
			}
			fmt.Println(i)

		}()
	}
	wg.Wait()
}

// 协程锁与释放
func Cond() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	for i := 0; i < 15; i++ {
		go func() {
			cond.L.Lock()
			cond.Wait()
			fmt.Println(".")
			cond.L.Unlock()
		}()

	}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 200)
		if i == 4 {
			cond.Signal() //单个协程
		}
		if i == 9 {
			cond.Broadcast() //全部协程
		}
	}
}
func Once() {
	var once sync.Once
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			fmt.Println("*")
			once.Do(func() {
				fmt.Println("只被打印一次")
			})
			defer wg.Done()
		}()
	}
	wg.Wait()
}
func Map() {
	var m sync.Map
	m.Store(1, "a")
	m.Store(2, "b")
	m.Store(3, "c")
	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}

// 锁
func RWMux() {

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

// sync.Mutex
func CountMutex() {
	var c int
	var mu sync.Mutex

	for i := 2; i < 100001; i++ {
		go func() {

			for j := 2; j < i; j++ {
				if i%j == 0 {
					return
				}
			}
			fmt.Printf("%v\t", i)
			mu.Lock()
			c++
			mu.Unlock()

		}()
	}
	time.Sleep(time.Second * 3)

	fmt.Println("总数", c)
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
