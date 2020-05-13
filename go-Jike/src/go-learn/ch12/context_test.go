package ch12

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancel(cxt context.Context) bool {
	select{
	case <- cxt.Done():
		return true
	default:
		return false
	}
}

func TestContext(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())//其中context.Background()用于获取父类，嵌套可获取子类
	for i:=0; i<=5; i++{
		go func(i int, ctx context.Context){
			for{
				if isCancel(ctx) {
					break
				}
				time.Sleep(time.Millisecond*10)
			}
			fmt.Println(i," is canceled!")
		}(i, ctx)
	}
	cancel()//用于取消当前context执行的task
	time.Sleep(time.Second*1)
}