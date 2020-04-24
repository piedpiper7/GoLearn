package ch13

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//适合于复用，降低复杂对象创建和GC代价
//但是生命周期依赖于GC，不适合做连接池，且协程安全有锁的开销

func TestSyncPool(t *testing.T){
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create new pool !")
			return 77
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(7758)
	runtime.GC()//GC 会清楚sync.pool中缓存的对象
	x := pool.Get().(int)
	fmt.Println(x)
}

func TestSyncPoolProcess(t *testing.T){
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new pool .")
			return 10
		},
	}

	pool.Put(110)
	pool.Put(120)
	pool.Put(130)

	wg := sync.WaitGroup{}
	for i:=0; i<10; i++ {
		wg.Add(1)
		go func(i int) {
			v := pool.Get().(int)
			fmt.Println(v)
			//time.Sleep(time.Millisecond*10)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
