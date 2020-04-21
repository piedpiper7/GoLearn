package main

import "fmt"

//递归实现阶乘
func main(){
	var i uint64 =1
	num := Factorial(i)
	fmt.Println(num)

	//递归实现Fibonacci数列
	for j:=1;j<20;j++{
		num := FibonacciSequence(j)
		fmt.Print(num," ")
	}

	fmt.Print("\n")

	for i:=1;i<20;i++ {
		y := fibonacci1(i)
		fmt.Print(y," ")
	}
}

func Factorial(n uint64) uint64 {

	if n<65{
		n = n*Factorial(n+1)
	}
	return n
}

func FibonacciSequence(i int) int {
	var result int
	if i>2{
		result = FibonacciSequence(i-1)+FibonacciSequence(i-2)
	} else{
		result = 1
	}
	return result
}

func fibonacci2(n int) (int,int) {
	if n < 2 {
		return 0,n
	}
	a,b := fibonacci2(n-1)
	return b,a+b
}


func fibonacci1(n int) int{
	_,b := fibonacci2(n)
	return b
}
