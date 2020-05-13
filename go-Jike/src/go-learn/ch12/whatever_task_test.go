package ch12

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string{
	time.Sleep(time.Millisecond*10)
	return fmt.Sprintf("%d task!", i)
}

func FirstResponse() string{
	taskNum := 10
	ch := make(chan string, taskNum)//防止协程泄露，及协程阻塞没有buffer的话
	for i:=0; i<taskNum; i++{
		go func(i int){
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch//一旦队列里有 就return
}

func AllResponse() string{
	numTask := 10
	ch := make(chan  string, numTask)
	for i:=0; i<= numTask; i++{
		go func(i int){
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	chStr := ""
	for j:=0; j<=numTask; j++{ //利用CSP机制实现全协程结束再返回，或者waitgroup
		chStr += <-ch + "\n"
	}
	return chStr
}

func TestRunTask(t *testing.T){
	t.Log("Before ",runtime.NumGoroutine())
	ret := FirstResponse()
	t.Log(ret)
	time.Sleep(time.Second*1)
	t.Log("After ",runtime.NumGoroutine())

	t.Log("Before ",runtime.NumGoroutine())
	ret1 := AllResponse()
	t.Log(ret1)
	time.Sleep(time.Second*1)
	t.Log("After ",runtime.NumGoroutine())
}