package main

import "fmt"

func main(){
	var a,b int = 1,2
	fmt.Println(a,b)
	testa()
	testb()
	_,c,d := return_value()
	fmt.Println(c,d)

}

func testa(){
	var a int
	a,b := 1,2
	fmt.Println(a,b)
}

func testb(){
	var a string = "abcde"
	b := a
	a = "cdef"
	fmt.Println(a,&a)
	fmt.Println(b,&b)
}
//返回值的函数
func return_value()(int,string,int){
	a, b, c := 1, "abc", 2
	return a,b,c
}
