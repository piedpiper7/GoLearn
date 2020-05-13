package ch10_inter

import (
	"fmt"
	"testing"
)

func DoSomething (p interface{}){
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}
	if j, ok := p.(string); ok {
		fmt.Println("String", j)
		return
	}
	fmt.Println("Unknown")
}

func DoSomething_a(p interface{}){
	switch v:=p.(type) {
	case int :
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown")
	}
}
func TestInterfaceAssert(t *testing.T){
	DoSomething(10)
	DoSomething("Gao")

	DoSomething_a(11)
	DoSomething_a("Hu")
}
