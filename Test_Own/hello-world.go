package main

import (
	"fmt"
	"os"
)

func main(){
	//fmt.Println(os.Args)
	if len(os.Args) > 3 {
		fmt.Println("Hello World!",os.Args)
	} else {
		fmt.Println("Hello World emm!")
	}
	os.Exit(999)
}
