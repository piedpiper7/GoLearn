package main

import (
	"fmt"
	"math"
)

func main(){

	var maxNum int
	var a, b string
	var x, y int = 100, 200
	maxNum = returnMax(5,19)
	a, b = changeA("are","you")
	fmt.Printf("The max number is %d\n",maxNum)
	fmt.Print(a," ",b)
	fmt.Printf("before a= %d b= %d\n", x, y)
	changeB(&x,&y)
	fmt.Printf("after a= %d b= %d\n", x, y)
    var num1 float64 = funcAA()
    fmt.Printf("square is %d", num1)
}

func returnMax(a ,b int) int {
	if a>b {
		return a
	} else {
		return b
	}
}

func changeA(a , b string) (string, string) {
	return b, a
}

func changeB(a *int, b *int)  {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

//函数作为实参
func funcAA() float64 {
	getSquareRoot := func (x float64) float64{
		return math.Sqrt(x)
	}
	return getSquareRoot(9)
}