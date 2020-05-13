package main

import (
	"fmt"
	"testing"
)

func testPrint(t *testing.T){
	//t.SkipNow()//跳过当前test，直接按照PASS执行下一test 必须写在首行
	res := Print1to20()
	fmt.Println("Hi")
	if res != 210{
		t.Errorf("Wrong result!")
	}
}

func testPrirt2(t *testing.T){
	res := Print1to20()
	fmt.Println("Hi")
	if res != 211{
		t.Errorf("Correct result!")
	}
}

//无法固定test执行顺序，但使用t.RUN可以执行subtests，控制test输出及顺序
//t.Run(para1, para2, para3) --para1(id),指向para2(函数)
//TestMain,初始化test，t.RUN可以调用其他test   runfirst

func TestAll(t *testing.T){
	t.Run("testPrint", testPrint)
	t.Run("testPrint2", testPrint)
}

func TestMain(m *testing.M){
	fmt.Println("Test Starting!")
	m.Run()
}

func BenchmarkAll(b *testing.B){
	for n:=0; n<b.N; n++{
		Print1to20()//执行b.N次，直到达到一个稳定值
	}
}
//go test -bench=.