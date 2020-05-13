package ch7

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringA(t *testing.T){
	s1 := "hello"
	t.Log(s1)
	t.Log(len(s1))
	s2 := "\xE0\xAA\xAB"
	t.Log(s2)
	t.Log(len(s2))

	s3 := []rune(s1)
	t.Log(s3)
	s4 := "中"
	s5 := []rune(s4)
	t.Log(s5)
	t.Log(len(s5))
	t.Log(&s4)
	t.Logf("unicode %x", s5[0])
	t.Logf("utf8 %x", s4)
}

func TestStringB(t *testing.T) {
   	s1 := "阿树宝贝"
   	for _, s := range s1{
   		t.Logf("%[1]c %[1]x", s)
	}
}

func TestStringDeal(t *testing.T){
	s1 := "A,B,C"
	parts := strings.Split(s1, ",")
	for _,v := range parts{
		t.Log(v)
	}
	s2 := strings.Join(parts,"#")
	t.Log(s2)

	s3 := 10
	result := strconv.Itoa(s3)
	t.Log("is string"+result)

	if i,err := strconv.Atoi("100"); err == nil{
		t.Log(i+10)
	}
}