package main

import (
	"fmt"
)

func main(){
	SlicePrint()
	twoDimensionArray()
}

func SlicePrint(){
	var SliceA = [] int {1,2,3}
	s := SliceA[:]
	s2 := SliceA[:2]
	s3 := SliceA[1:2]
	var result [] int
    copy(result,ManyToOne(s,s2,s3))
	result2 := make([]int, len(s)+len(s2)+len(s3))
	copy(result2,ManyToOne(s,s2,s3))
	fmt.Println("!!!",result)//不可以
	fmt.Println("!!!",ManyToOne(s,s2,s3))
	fmt.Println("!!!",result2)

	for i:=range s{
		fmt.Println(s[i])
	}
}

func twoDimensionArray(){
	var arrayB = [][]int {{0,0},{2},{3},{3,4}}
	fmt.Println(arrayB)
	for i:=0;i<len(arrayB);i++{
		for j:=0;j<len(arrayB[i]);j++{//len(arrayB[i])等价于取内部的{},长度就是内部{}内元素个数
			fmt.Print(arrayB[i][j],"\n")
		}
	}
}

func ManyToOne(a []int, b []int, c []int) []int {
	return append(append(a, b...), c...)//必须加...
}