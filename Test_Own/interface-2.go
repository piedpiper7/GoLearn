package main

import "fmt"

type XPhone interface {
	call() string
}

type IPhone1 struct {
	version string
}

type Huawei1 struct {
	type1 string
}

type Xiaomi struct {
	money string
}

func (iphone IPhone1) call() string{
	return "i am iphone"+iphone.version
}

func (huawei Huawei1) call() string{
	return "i am huawei"+huawei.type1
}

func (xiaomi Xiaomi) call() string{
	return "i am xiaomi " + xiaomi.money
}

func printcall(p XPhone){
	fmt.Println(p.call()+",i can call you!")
}

func main(){
	var huawei = Huawei1{"aaa"}
	var iphonex = IPhone1{"11 pro"}

	//var xiaomia XPhone
	//xiaomia = new(Xiaomi)
	var xiaomia = Xiaomi{"2000"}
	t := xiaomia.call()
	fmt.Println(t)

	printcall(huawei)
	printcall(iphonex)
}