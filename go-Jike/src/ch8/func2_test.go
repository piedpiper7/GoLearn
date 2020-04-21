package ch8

import (
	"fmt"
	"testing"
)

func Sum(ops ... int) int {
	result := 0
	for _, op := range ops{
		result += op
	}
	return result
}

func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3))
	t.Log(Sum(1,2,3,4))

}

func Clear(){
	fmt.Println("Clear Resources.")
}

func TestDefer(t *testing.T){
	defer Clear()//defer可用于释放资源，释放锁
	fmt.Println("START!")
	panic("err")//存在panic，则之后的code都不会执行，但是defer会执行
 }