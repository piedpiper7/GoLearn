package main

import (
	"fmt"
	"strings"
)

var c bool

func main(){
	var a float32 = 1.2
	var b int
	b = 10
	c = true
	d := "abcde"
	fmt.Println("abcd"+"@@")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	test1()

	space_out()
}

func test1(){
	c = false
	fmt.Println(c)
}

//去除空格 和 换行
func space_out(){
	var str string = "hello\n world\n!"
	fmt.Println(str)
	fmt.Println(strings.Replace(str,"\n","",-1))
	fmt.Println(strings.Replace(str," ","",-1))
	str2 := strings.Replace(str,"\n","",-1)
	fmt.Println(str2)

}