package ch5

import "testing"

func TestArray(t *testing.T){
	var arry1 = [3] int {1,2,3}
	arry2 := [...] int {2,4,6,8}

	arry3 := [3][2] int {{1,2},{2,3},{3,4}}

	for i := 0; i < len(arry1); i++{
		t.Log(arry1[i])
	}

	for _, e := range arry2{
		t.Log(e)
	}

	t.Log(arry3[1][0])
}

func TestArray2(t *testing.T){
	var array1 = [...] int {1,2,3,4,5}
	t.Log(array1[2:])
	t.Log(array1[2:len(array1)])
	t.Log(array1[1:3])
}

func TestSlice(t *testing.T){
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1, 2)
	t.Log(len(s0), cap(s0))
	t.Log(s0[1])

	s1 := make([]int, 3, 5)
	t.Log(len(s1), cap(s1))
	s1 = append(s1,3)
	t.Log(s1[0], s1[1], s1[2], s1[3])

}

func TestSlice2(t *testing.T){
	var s2  []int
	for i := 1; i <= 17; i++{
		s2 = append(s2, i)
		t.Log(len(s2), cap(s2))
	}
}

func TestSliceShare(t *testing.T){
	Q1 := []string {"Jua", "Feb", "May", "Apr", "June", "July", "Aug", "Sep", "Nov", "Oct", "Dec"}
	Spring := Q1[0:4]
	Summer := Q1[3:6]
	t.Log(Spring, len(Spring), cap(Spring))
	t.Log(Summer, len(Summer), cap(Summer))

	Summer[2] = "unknown"
	Summer[0] = "unknown"

	t.Log(Summer)

	t.Log(Spring)
}

func TestComparing(t *testing.T){
	a := [...]int{1,2,3}
	b := [...]int{1,2,3}
	if a == b {
		t.Log("true")
	} else {
		t.Log("false")
	}
}