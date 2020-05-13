package ch10_inter

import (
	"fmt"
	"testing"
)

type Code string

type programmer interface {
	WriteHello() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHello() Code {
	return "fmt.Println(\"Hello World !\")"
}

type JavaProgrammer struct {

}

func (g *JavaProgrammer) WriteHello() Code {
	return "System.out.println(\"Hello World ! \")"
}

func WriteFirstHello(p programmer){
	fmt.Printf("%T %v\n",p, p.WriteHello())//%T是输出类型
}

func TestInterface(t *testing.T){
	//var p programmer
	//p = new(GoProgrammer)
	//var q programmer
	//q = new(JavaProgrammer)
	//t.Log(p.WriteHello())
	//t.Log(q.WriteHello())
	p := new(GoProgrammer)
	WriteFirstHello(p)
	q := new(JavaProgrammer)
	WriteFirstHello(q)

}