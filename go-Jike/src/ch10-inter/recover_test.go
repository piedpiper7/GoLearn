package ch10_inter

import (
	"errors"
	"fmt"
	"testing"
)
//不建议把异常直接打印成日志，会形成僵尸程序
func TestPanicRecover(t *testing.T){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("recover error ", err)
		}
	}()

	panic(errors.New("Go out "))
}
