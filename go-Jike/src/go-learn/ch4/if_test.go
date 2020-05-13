package ch4

import "testing"

func TestIF(t *testing.T){

	n := 1

	for n<5 {
		t.Log(n)
		n++
	}

	for i := 1; i < 5; i++{
		t.Log(i)
	}
}
