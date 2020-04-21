package Fib

import "testing"

func TestFibList(t *testing.T){
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T){
	a := 1
	b := 2
	b, a = a, b
	t.Log(a,b)
}


