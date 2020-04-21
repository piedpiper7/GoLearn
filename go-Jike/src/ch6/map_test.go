package ch6

import "testing"

func TestMap_a(t *testing.T){
	map1 := map[string] int {"one":1,"two":2}
	map2 := map[int] int {}
	map2[1] = 1
	map3 := make(map[string] int, 10)
	t.Log(map1)
	t.Log(map2)
	t.Log(map3)
    t.Log(len(map3))
}


func TestMap_b(t *testing.T){
	map1 := map[int] int {}
	t.Log(map1[1])
	map1[2] = 0
	if v, ok := map1[2]; ok{
		t.Logf("is existing, value is %d ", v)
	} else {
		t.Log("is not existing")
	}

}

func TestMap_c(t *testing.T){
	map1 := map[int] int {1:2, 2:4, 3:6}

	for k, v := range map1{
		t.Log(k, v)
	}
}