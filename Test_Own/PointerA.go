package main

import "fmt"

func main(){
	pointA()
	arrayPtr()
	pptrTest()
}

func pointA(){
	var a1 int = 7
	var add *int
	var null *int
	add = &a1
	fmt.Println(add)
	fmt.Println(*add)
	fmt.Println(null)
	var a,b int = 3,4
	swapPtr(&a,&b)
	fmt.Printf("a=%d,b=%d\n",a,b)
	swapA(&a,&b)
	fmt.Printf("a=%d,b=%d\n",a,b)
}

const max int = 4

func arrayPtr(){
	var arrayA = [] int {1,2,3,4}
	var prtA [max] *int
	for i:=0;i<len(arrayA);i++{
		prtA[i] = &arrayA[i]
	}
	for i:=0;i<max;i++{
		fmt.Println(*prtA[i])
	}
}

func pptrTest(){
	var ptr *int
	var pptr **int
	var x int = 10
	ptr = &x
	pptr = &ptr
	fmt.Println(*pptr)
}

//指针作为函数参数
func swapPtr(x *int, y *int){
	var change int
	change = *x
	*x = *y
	*y = change
}
func swapA(x *int, y *int){
	*x,*y = *y,*x
}