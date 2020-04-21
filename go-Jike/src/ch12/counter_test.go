package ch12

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T){
	counter := 0
	for i:=0; i<5000; i++{
		go func(){
			counter++
		}()
	}
	time.Sleep(time.Second*1)
	t.Log(counter)
}

func TestCounterSafe(t *testing.T){
	counter := 0
	var mut sync.Mutex
	for i:=0; i<5000; i++{
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second*1)
	fmt.Println("Safe counter ", counter)
}

func TestCounterWaitGroup(t *testing.T){
	counter := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	for i:=0; i<5000; i++{
		wg.Add(1)
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}