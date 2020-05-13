package ch12

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type singleObject struct {

}

var once sync.Once
var singleInstance *singleObject

func singleOnce() *singleObject{
	once.Do(func() {
		fmt.Println("Object created!")
		singleInstance = new(singleObject)
	})
	return singleInstance
}

func TestSingleOnce(t *testing.T){
	var wg sync.WaitGroup
	for i:=1; i<=7; i++{
		wg.Add(1)
		go func(){
			obj := singleOnce()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}