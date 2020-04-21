package main

import "fmt"

func main(){
	YanghuiOut(13)
}

func YanghuiOut(x int){
	for i:=0; i<x ; i++{
		num := 1
		for j:=0; j<x-i; j++{
			fmt.Print(" ")
		}
		for j:=0; j<=i; j++{
			fmt.Printf("%5d", num)
			num = num*(i-j)/(j+1)
		}
		fmt.Println("\n")
	}
}
