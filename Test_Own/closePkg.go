package main
//匿名函数，作为闭包，免声明
import (
	"fmt"
)

func main(){
	nextNum := getSequence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	nextNum1 := getSequence()
	fmt.Println(nextNum1())
	fmt.Println(nextNum1())

	addFunc := add(1,2)
	fmt.Println(addFunc())
	fmt.Println(addFunc())

	addBnum := addB(3,4)
	fmt.Println(addBnum(4,6))
	fmt.Println(addBnum(4,6))


}

func getSequence() func() int{
	i:=0
	return func() int{
		i+=1
		return i
	}
}

func add(x, y int) func() (int, int){
	i:=1
	return func() (int ,int){
		i++
		return i,x+y
	}
}

func addB(x,y int) func(a,b int)(int,int,int){
	i:=1
	return func(a, b int) (int, int,int) {
		i++
		return i,x+y,a+b
	}
}