package ch14

import (
	"fmt"
	"testing"
)

func TestSquaresUnit(t *testing.T){
	inputs := [...] int {1,2,3}
	//exceptions := [...] int {1,4,9}
	exceptions := [...] int {4,5,6}
	for i:=0; i< len(inputs); i++{
		ret := Square(inputs[i])
		if ret != exceptions[i] {
			t.Errorf("Input is %d, Exceptions is %d, But acctual is %d",inputs[i], exceptions[i], ret)
		}
	}
}


//difference between error and fetal

func TestError(t *testing.T){
	fmt.Println("start")
	t.Error("Error")
	fmt.Println("End")
}

func TestFetal(t *testing.T){//中断当前测试，继续执行其他的测试，所以不会输出END
	fmt.Println("start")
	t.Fatal("Error")
	fmt.Println("End")
}