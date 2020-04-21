package main

import (
	"fmt"
)

func main(){
	var a bool = true
	var b bool = false
	fmt.Println(a&&b)
	fmt.Println(a||b)
	fmt.Println(!(a&&b))
	fmt.Println(!a)
	bit_calc()
	test_b()
}

func bit_calc(){
	var a  uint = 60//111100
	var b uint = 13 // 1101
	var c uint = 0
	c = a & b
	fmt.Println(c)//12
	c = a | b
	fmt.Println(c)
	c = a ^ b
	fmt.Println(c)
	c = a << 2
	fmt.Println(c)
	c= a >> 2
	fmt.Println(c)
}

func test_b(){
	var a int = 4
	var ptr *int
	ptr = &a
	fmt.Println(a)
	fmt.Println(ptr)
}