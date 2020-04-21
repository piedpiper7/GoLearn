package ch6

import "testing"

func TestFac_a(t *testing.T){
	map1 := map[int] func(op int) int {}
	map1[1] = func(op int) int {return op}
	map1[2] = func(op int) int {return op*op}
	map1[3] = func(op int) int {return op*op*op}
	t.Log(map1[1](3), map1[2](3), map1[3](3))
}

func TestSet(t *testing.T){
	mapSet := map [int] bool {}
	mapSet[1] = true

	mapSet[2] = true
	n2 := 2
	if mapSet[n2] {
		t.Logf("%d is exsting", n2)
	}	else {
		t.Logf("%d is not exist", n2)
	}
	delete(mapSet,1)

	n := 1
	if mapSet[n] {
		t.Logf("%d is exsting", n)
	}	else {
		t.Logf("%d is not exist", n)
	}

	t.Log(len(mapSet))
}
