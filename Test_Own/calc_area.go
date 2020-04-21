package main

import ("fmt"
"unsafe"
)

const pai = 3

func main(){
	var a int
	b := 2
	a = pai*b*b
	fmt.Println(a)
	test_a()
	iota_test()
}

func test_a(){
	const (a  = "abcde"
	b = len(a)
	c = unsafe.Sizeof(a)
	)
	fmt.Println(a,b,c)
}

func iota_test(){
	const (a  = iota
	b
	c)
	fmt.Println(a,b,c)
	const (d  = 1<<iota //001
	e = 3<< iota //110 = 6
	f//1100 = 12
	g// 11000= 24
	)
    fmt.Println(d,e,f,g)
	const (h  = 2<<iota //10
		i = 4<< iota //10000
		j //100000
		k// 1000000
		 )
		fmt.Println(h,i,j,k)
}
