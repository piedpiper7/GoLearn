package ch12

import (
	"fmt"
	"testing"
	"time"
)

func cancelCheck(cancleCh chan struct{}) bool{
	select  {
	case <- cancleCh:
		return true
	default:
		return false
	}
}

func channel_1(cancelCh chan struct{}){
	cancelCh <- struct{}{}
}

func channel_2(cancelCh chan struct{}){
	close(cancelCh)
}

func TestCancelChannel(t *testing.T){
	cancelChan := make(chan struct{}, 0)
	for i:=0; i<5; i++{
		go func(i int, cancelCh chan struct{}){
			for {
				if cancelCheck(cancelCh) {
				break
			}
			time.Sleep(time.Millisecond*5)
			}
			fmt.Println(i, "is canceled!")
		}(i, cancelChan)
	}
	//channel_1(cancelChan)
	channel_2(cancelChan)//广播机制，所有的协程都会收到close
	time.Sleep(time.Second*1)
}