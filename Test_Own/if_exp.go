package main

import "fmt"

func main(){
	var a = 1;
	if a > 1 {
		fmt.Println("a<1")} else {
			fmt.Println("a<1hhh")}
	fmt.Println(a)
	if_cascade()
	switch_a(100)
	switch_a(10)
	select_a()
}

func if_cascade() {
	var a int = 200
	var b int = 200
	if a == 100 {
		if b < 100 {
			fmt.Println("a是100，且b小于100")
		} else {
			fmt.Println("a is 100 and b more than 100")
		}
	} else {
		fmt.Println("nothing interesting!")
	}
}

func switch_a(num int){
	//var num int = 100
	var level string
	switch num{
	case 100: level = "beauty"
	case 60: level = "just so so"
	case 10:level = "ugly"
	default:level = "i love you"
	}
	fmt.Printf("your grade is %d and your are %s\n",num, level)
}

func select_a(){

	var c1, c2, c3 chan int
	var i1, i2 int
	select{
	case i1 = <-c1:
		fmt.Printf("recivied", i1, "from c1\n")
	case c2<- i2:
		fmt.Printf("sent" ,i2, "to c2\n")
	case i3, ok:= (<-c3):
		if ok {
			fmt.Printf("reciviced", i3, "from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communicate\n")
	}



}