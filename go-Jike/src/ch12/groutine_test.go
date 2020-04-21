package ch12

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T){
	for i := 0; i <= 20; i++{
		go func (i int){
			fmt.Println("value is ", i)
		}(i)
		time.Sleep(time.Millisecond*50)
	}
}
