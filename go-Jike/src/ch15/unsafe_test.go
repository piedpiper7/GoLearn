package ch15

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

//不安全的转换
func TestUnsafe(t *testing.T){
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))//利用unsafe取出指针，再强制转换成浮点指针
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

//合理的转换

type MyInt int

func TestSafeSlice(t *testing.T){
	s1 := [] int {1,2,3,4}
	s2 := *(*[] MyInt)(unsafe.Pointer(&s1))
	t.Log(s2)
}

//原子类型操作
//并发读写时 需要线程安全 在写完之后 进行原子操作，指向新的内容

func TestAtomic(t *testing.T){
	var shareBufPtr unsafe.Pointer
	dataWrite := func(){
		data := []int{}
		for i:=0; i<100; i++{
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))//用原子操作将data置于共享指针
	}
	readData := func(){
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(&data,*(*[]int)(data))
	}
	var wg sync.WaitGroup
    dataWrite()
	//读和写的协程各十个
	for i:=0; i<10; i++{
		wg.Add(1)
		go func(){
			for j:=0; j<10; j++{
				dataWrite()
				time.Sleep(time.Millisecond*10)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func(){
			for j:=0; j<10; j++{
				readData()
				time.Sleep(time.Millisecond*10)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
