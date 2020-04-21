package main

import (
	"fmt"
)

func main(){
    arrayA()
    deleteArray()
}

func arrayA(){
	var nameList = [...] string {"i", "love", "you"}
	for i:=0; i < len(nameList); i++{
		fmt.Printf("%s ", nameList[i])
	}
}

func deleteArray(){
	var arrayA = [] int {1, 2, 3, 4}
	for j:= 0; j<len(arrayA);j++{
		arrayA = append(arrayA[:j])
	}
	for i:=0; i < len(arrayA); i++{
		fmt.Printf(" %d", arrayA[i])
	}
}

