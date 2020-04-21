package main
//???????????
import "fmt"

type Fruit interface {
	GetName() string
	SetName() string
}

type Apple struct {
	name string
}

func (a *Apple) GetName(){

}

func (a *Apple) SetName(){

}

func main(){
	//var a Fruit
	a := Apple{"red apple"}
	fmt.Println(a)
	//a:=Apple{"green apple"}
	//fmt.Println(a)
}