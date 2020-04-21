package main

import "fmt"

func main(){
	var ptr1 *int
	var a int = 2
	var ptr2 *int
	//ptr2 = &a
	ptr1 = &a
	switch  {
	case ptr1 != nil:
		fmt.Printf("ptr1 is not nil\n")
		fallthrough
	case ptr1 == nil:
		fmt.Printf("ptr1 is nil\n")
		//fallthrough
	case ptr2 != nil:
		fmt.Printf("ptr2 is not nil\n")
	default:
		fmt.Printf("ptr2 is nil\n")
	}
}
