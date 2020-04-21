package main

import "fmt"

func main (){
	sliceUse()
	AppendCopy()
}

func sliceUse(){
	var arrayA = [4] int {1,2,3,4}
	//定义切片s是arrayA的引用
	s := arrayA[:]
	fmt.Println(s)
	fmt.Println("******************")
	lenCap(s)
	//将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	s1 := arrayA[1:3]
	fmt.Println(s1)
	fmt.Println("******************")
	lenCap(s1)
	//默认 endIndex 时将表示一直到arr的最后一个元素
	s2 := arrayA[:3]
	fmt.Println(s2)
	fmt.Println("******************")
	lenCap(s2)
	//通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
	s3 := make([] int,3,7)
	fmt.Println(s3)
	lenCap(s3)
}

//len() & cap()
func lenCap(x []int){
	fmt.Println(x,len(x),cap(x))
}

//append() & copy()
func AppendCopy(){
	var sliceA []int
	lenCap(sliceA)
	sliceA=append(sliceA, 1,2,3)
	lenCap(sliceA)
	sliceB := make([]int, len(sliceA),cap(sliceA)*2)
	lenCap(sliceB)
	copy(sliceB,sliceA)
	lenCap(sliceB)

}