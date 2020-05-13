package ch9_OO

import (
	"fmt"
	"testing"
	"unsafe"
)

type employee struct{
	Id string
	Name string
	Age int
}

func TestCreateEEmployeeObj(t *testing.T){
	e := employee{"123","GJN",18}
	e1 := employee{Id:"124", Age:17}
	e2 := new(employee)
	e2.Id = "125"
	e2.Name = "Xi"
	e2.Age = 17
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Name)
	t.Log(e2)
	t.Logf("e2 is %T", e2)
}

func (e *employee) String() string{//不使用指针会导致结构复制，使得空间开销增加
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-NAME:%s-AGE:%d",e.Id,e.Name,e.Age)
}

func TestEmployeeString(t *testing.T){
	e := employee{"123", "Bob",21}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}