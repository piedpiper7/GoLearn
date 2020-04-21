package main

import "fmt"

func main(){
	SADiff()
	appendSlice()
}

func SADiff(){
	var arrayA = [3] int {1,2,3}
	var SliceA = [] int {4,5,6}
	ChangeArr(arrayA)
	fmt.Println(arrayA)
	ChangeSlice(SliceA)
	changeArrPtr(&arrayA)
	fmt.Println(SliceA)
	fmt.Println(arrayA)

}

//slice传入的是引用
func ChangeArr(arr [3] int){
	for i:=0;i<len(arr);i++{
		arr[i] = arr[i] +1
	}
}

func ChangeSlice(slice [] int){
	for i:=0;i<len(slice);i++{
		slice[i] = slice[i] +1
	}
}

//使用切片等价于数组使用指针
func changeArrPtr(arr *[3]int){
	for i:=0;i<len(arr);i++{
		arr[i] = arr[i] +1
	}
}

//对切片进行扩容
func appendSlice(){
	var sliceA = [] int {1,3,7}
	lenCap1(sliceA)
	sliceA=append(sliceA, 2,3,1,4,5)
	lenCap1(sliceA)
	/*sliceA=append(sliceA, 4,5)
	lenCap1(sliceA)
	sliceA=append(sliceA, 6)
	lenCap1(sliceA)
	sliceA=append(sliceA, 7,8,9)
	lenCap1(sliceA)
	sliceA=append(sliceA, 7,9,10,11,12,13,14)
	lenCap1(sliceA)
	sliceA=append(sliceA, 7,8)
	lenCap1(sliceA)
	sliceA=append(sliceA, 7,8,9,10)
	lenCap1(sliceA)
	sliceA=append(sliceA, 6)
	lenCap1(sliceA)*/
}

func lenCap1(x []int){
	fmt.Println(x,"len:",len(x),"cap=",cap(x))
}