package _const

import "testing"

const (
	Monday = iota +1
	Tuesday
	Wednesday
)

const (
	//连续常量 位常量
	read = 1 << iota
	write
	execute
)

func TestConst(t *testing.T){
	t.Log(Monday, Wednesday)
}

func TestConst2(t *testing.T){
	a := 7
	b := 8
	a = a &^ read// 0111  1  0110
	b = b &^ execute/// 1000 0011 1000
	t.Log(b)
	t.Log(a)
	t.Log(a&read == read, a&write == write, a&execute == execute)
	t.Log(b&read == read, b&write == write, b&execute == execute)

}