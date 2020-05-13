package ch10_inter

import (
	"errors"
	"fmt"
	"testing"
)

func fibonacci(n int) ([]int, error){
	if n < 2 || n >100 {
		return nil, errors.New("n must in [2,100]")
	}
	FibList := [] int {1,1}
	for i := 2; i <= n; i++ {
		FibList = append(FibList, FibList[i-1] + FibList[i-2])
 	}
	return  FibList, nil
}

func TestFibonacciErr(t * testing.T){
	if v, err := fibonacci(-10); err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}