
package main

import "fmt"

type music struct {
	title string
	auther string
	Mtype string
}

func main (){
	fmt.Println(music{"Five Hang","Jianning","jazzs"})
	fmt.Println(music{title:"AAA",Mtype:"Happy"})
	RequestStruct()
}

//访问结构体成员
func RequestStruct(){
	var music1 music
	var music2 music
	music1.title = "As long as you love me"
	music1.auther = "Justicing"
	music2.title = "he's not that into you"
	music2.Mtype = "sad"
	fmt.Printf("music1 is %s by %s\n", music1.title,music1.auther)
	fmt.Printf("music2 is %s type is %s\n",music2.title,music2.Mtype)
	StructPrint(music1)
	StructPrint(music2)
	StructPrint1(&music1)
}

//作为参数传递给函数
func StructPrint(x music){
	changStruct(&x)
	fmt.Println(x.title)
	fmt.Println(x.auther)
	fmt.Println(x.Mtype)
}

//作为指针
func StructPrint1(x *music){
	fmt.Println(x.title)
	fmt.Println(x.auther)
	fmt.Println(x.Mtype)
}

//改变结构体的内容
func changStruct(x *music){
	x.title = "As short as you love me"
}