package main

import "fmt"

func main(){
	var result float32
	arrayB := [] int {1,3,4,7}
	result = Calculation(arrayB)
	fmt.Println(result)
	fmt.Println(len(arrayB))
}
func Calculation(arrayA [] int) float32{
	lenArr := len(arrayA)
	num := 0
	for i:=0;i<lenArr;i++{
		num += arrayA[i]
	}
	for i:=0; i< lenArr; i++{
		arrayA = append(arrayA[:i])
	}
	return float32(num)/float32(lenArr)
}