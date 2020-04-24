package ch12

import (
	"fmt"
	"testing"
	"time"
)

//communication sequensetinal process

func service() string {
	time.Sleep(time.Millisecond*150)
	return "Done task1"
}

func otherService(){
	fmt.Println("Wroking on other!")
	time.Sleep(time.Millisecond*60)
	fmt.Println( "Done task2")
}

func AsynServices() chan string{
	retCH := make(chan string, 1)//buffer channel's capacity
	go func(){
		ret := service()
		fmt.Println("return result")
		retCH <- ret
		fmt.Println("service in channel")
	}()
	return retCH
}
//func TestCsp(t *testing.T){
//	t.Log(service())
//	otherService()
//}
func TestAsynServices(t *testing.T){
	retCh := AsynServices()
	otherService()
	t.Log(<-retCh)
	time.Sleep(time.Second*1)
}

func TestServices(t *testing.T){
	select {
	case retCh := <-AsynServices():
		t.Log(retCh)
	case <-time.After(time.Millisecond*100):
		t.Error("timeout")
	}
}