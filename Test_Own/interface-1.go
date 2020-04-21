package main

//接口方法传参

import "fmt"

type Phone interface {
	call(param int) string
	takephoto()
	takephoto2() string
}

//可以定义结构体的类型
type Huawei struct {
	num int
	type1 string

}

func (huawei Huawei) call(param int) string{
	fmt.Println("i am huawei!", Huawei{num:20})
	return "damon"
}

func (huawei Huawei) takephoto2() string{
    return "aaa" + huawei.type1
}

func (huawei Huawei) takephoto() {
	fmt.Println("i can take photo!")
}

func main(){
	var phone Phone
	//实例化接口
	phone = new(Huawei)
	phone.takephoto()
	r := phone.call(50)
	fmt.Println(r)
	x := phone.takephoto2()
	var x1 Phone
	x1 = Huawei{type1:"new one"}
	fmt.Println(x1)

	fmt.Println(x)
}
