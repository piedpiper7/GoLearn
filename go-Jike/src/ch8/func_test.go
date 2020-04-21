package ch8

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type IntConv func(op int) int

func returnMultValues() (int, int){
	return rand.Intn(10), rand.Intn(5)
}

func TimeSpend( inner IntConv) IntConv {//自定义类型
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend :", time.Since(start).Seconds())
		return ret
	}
}

func TimeSlow(op int) int {
	time.Sleep(time.Second*1)
	return op
}
func TestFunction(t *testing.T){
	a, b := returnMultValues()
	t.Log(a, b)
	spend := TimeSpend(TimeSlow)
	t.Log(spend(10))
}