package main

import "fmt"

func main(){
	break_t()
	continue_a()
}

func break_t(){

	for i:=1;i<=100;i++{
		fmt.Printf("this is %d\n",i)
		if i>=5 {
			break
		}
	}
}

func continue_a(){
	for i:=1;i<=10;i++{
		fmt.Printf("this is %d\n",i)
		if i%2 ==0 {
			fmt.Printf("%d also is hh\n",i)
			continue
		}
		fmt.Printf("_________________________\n")
	}
}
