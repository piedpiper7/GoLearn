
package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {

}

type IPhone struct{

}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am NokiaPhone,hello")
}

func (iPhone IPhone) call() {
	fmt.Println("i am NokiaPhone,hello")
}

func main (){
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}