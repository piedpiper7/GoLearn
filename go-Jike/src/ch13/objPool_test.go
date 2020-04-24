package ch13

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T){
	pool := NewObjPool(10)
	if err := pool.ReleaseObj(&ReusableObj{}); err != nil{
		t.Error(err)
	}//满了还放
	for i:=0; i<11; i++{
		if v, err:= pool.GetObj(time.Second*1); err != nil{
			t.Error(err)
		} else {//如果不释放，只加在第11个会因为等待超时而timeout
			fmt.Printf("%T\n", v)
			//if err := pool.ReleaseObj(v); err != nil {
			//	t.Error(err)
			//}
		}
	}
}
