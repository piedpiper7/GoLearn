package main

import (
	"fmt"
)

type DivideError1 struct {
	ErrType int
	var1 int
	var2 int
}

func (divideError1 DivideError1) Error() string{
	if 0 == divideError1.ErrType{
		return "divider is zero"
	} else {
		return "other wrong"
	}
}

func Divide1(a int, b int) (int, *DivideError1){
	if b==0 {
		return 0,&DivideError1{0,a,b}
	} else{
		return a/b,nil
	}
}

func main(){
	v,r :=Divide1(100,2)
	if nil!=r{
		fmt.Println("(1)fail:",r)
	}else{
		fmt.Println("(1)succeed:",v)
	}
	// 错误调用
	v,r =Divide1(100,0)
	if nil!=r{
		fmt.Println("(2)fail:",r)
	}else{
		fmt.Println("(2)succeed:",v)
	}

}