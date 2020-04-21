package main

import "fmt"

func main(){
	num := [] int {}
	for i:=0; i<10; i++{
		num = arrayAppend(num)//每一次入参为上一次获得数组
		fmt.Println(num)
	}
}

func arrayAppend(inArrary [] int) [] int{
	var out []int
	out = append(out,1)
    inArrlen := len(inArrary)
    if inArrlen == 0 {
    	return out
	}
    for i :=0; i<inArrlen-1 ;i++{
    	out = append(out, inArrary[i] + inArrary[i+1])
	}
	out = append(out,1)
    return out
}
