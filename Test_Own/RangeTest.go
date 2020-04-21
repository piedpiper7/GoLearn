package main


import (
	"fmt"
	"os"
)

func main(){
	rangeList()
}

func rangeList(){
	var sliceA  = [] int {1,2,3}
	var sum int
	for _,num := range sliceA{
		sum+=num
	}
	fmt.Println(sum)

	for i:= range sliceA{
		fmt.Println(sliceA[i])
	}

	kvs := map[string] string {"you":"ZXY","me":"GJN"}
	for k,v := range kvs{
		fmt.Println(k,v)
	}
	//Unicode枚举
	for i,c := range "ZXYWXHN"{
		fmt.Print(i,c," ")
	}
	var str = "ZXYWXHN"
	var s *string
	s =&str
	fmt.Println(*s)

	for _,i := range os.Args{
		fmt.Print(i," ")
	}
}