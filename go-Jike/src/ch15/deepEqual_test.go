package ch15

import (
	"fmt"
	"reflect"
	"testing"
)

//map和切片
func TestDeepEqual(t *testing.T){
	map1 := map [int] string {1:"one", 2:"two", 3:"three"}
	map2 := map [int] string {1:"one", 2:"two", 3:"three"}
	fmt.Println(reflect.DeepEqual(map1,map2))

	s1 := [] int {1,2,3}
	s2 := [] int {1,2,3}
	s3 := [] int {1,4,3}
	t.Log("s1 == s2 ?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3 ?", reflect.DeepEqual(s1, s3))
}
