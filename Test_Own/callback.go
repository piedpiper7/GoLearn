package main

import "fmt"

type ft func(int) int

func main(){
	oneCallBack(1,twoCallBack)
	oneCallBack(2,func(x int) int {
		fmt.Printf("callback is %d", x)
		return x
	})
}

func oneCallBack(x int, f ft){
	f(x)
}

func twoCallBack(x int) int {
	fmt.Printf("callback is %d\n", x)
	return x
}