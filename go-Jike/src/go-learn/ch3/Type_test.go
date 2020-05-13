package ch3

import (
	"testing"
)

type MyInt int64

func TestType(t *testing.T){
	var a int32 = 1
	var b int64
	b = int64(a)

	t.Log(b)

	var c MyInt
	c = MyInt(b)
	t.Log(c)
}


func TestPoint(t *testing.T){
	var a = 1
	aPtr := &a
	t.Log(a," ",aPtr)
	t.Logf("%T %T",a ,aPtr)

	var b string
	t.Log("*",b ,"*")
	t.Log(len(b))
	if b == "" {
		t.Log("Yes!")
	}
}