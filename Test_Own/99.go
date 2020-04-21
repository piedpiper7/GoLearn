package main

import "fmt"

func main(){
 calc(8,9)
 goto99()
}

func calc(a int, b int){
	for i:= 1;i<=a; i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%d  ",i,j,i*j)
		}
		fmt.Printf("\n")
	}
}

func goto99(){
	for m:=1;m<10;m++{
		n:=1
		loop: if n<=m{
			fmt.Printf("%d*%d=%d  ",m,n,m*n)
			n++
			goto loop
		} else {
			fmt.Printf("\n")
		}
		//n++
	}
}