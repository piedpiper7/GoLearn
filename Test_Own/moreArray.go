package main

import "fmt"

func main(){
	PrintArray()
}

func PrintArray(){
	arrayA := [3][4] int{{1,2,3,4},
	{5,6,7,8},
	{9,10,11,12},
	}
	for i:=0;i<3; i++{
		for j:=0; j<4; j++{
			fmt.Print(arrayA[i][j]," ")
		}
		fmt.Println("\n")
	}
}