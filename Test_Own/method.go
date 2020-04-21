package main

import "fmt"

type Circle struct {
	radius float64
}//定义结构体

func main(){
	var c1 Circle
	c1.radius = 10.00//c1是对象，属性是radius
	fmt.Println("area is ",c1.getArea())
}

func (c Circle) getArea() float64{//circle类型的方法
	return c.radius*c.radius*3.14
}
