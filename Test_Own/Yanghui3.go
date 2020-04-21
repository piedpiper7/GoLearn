package main

import "fmt"

func main(){
	yangHuiOut(13)
}

func yangHuiOut( size int){
	arrayA := []int{}
	for i:= 0; i< size ; i++ {
		lenArr := len(arrayA)
		if lenArr == 0 {
			arrayA = append(arrayA,1)
		} else {
			arrayB := []int{1}
			for i:=0; i< lenArr-1; i++{
				arrayB = append(arrayB,arrayA[i]+arrayA[i+1])
			}
			arrayB = append(arrayB,1)
			arrayA = arrayB
		}
		fmt.Println(arrayA)
	}
}
