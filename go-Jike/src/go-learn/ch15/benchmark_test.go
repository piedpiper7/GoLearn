package ch15

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStrByAdd(t *testing.T){
	assert := assert.New(t)
	elems := [] string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems{
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStrByBuffer(t *testing.T){
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := [] string{"1", "2", "3", "4", "5"}
	for _,elem := range elems{
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}

func BenchmarkTest(b *testing.B) {
	elems := [] string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkTest2(b *testing.B){

	elems := [] string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i:=0; i<b.N; i++{
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
}

//benchmark的使用规则
//func 以Benchmark开头
//与性能测试无关的代码
//b.resetTimer()
//for i:=0; i< b.N; i++{
//测试代码
//}
//b.stopTomer()