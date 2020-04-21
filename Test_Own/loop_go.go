package main

import "fmt"

func main(){
	for_loop()
	prime_number()
	prime_num_b()
}

func for_loop(){
	var a int = 15
	var b int

	numbers := [6] int {1,2,3,5}//s数组的长度是6，即从0-5

	for b:=0; b<10; b++{
		fmt.Printf("b的值是%d\n", b)
	}
	for b<a {
		b++
		fmt.Printf("b的另一个值是%d\n",b)
	}

	for i,x := range numbers{
		fmt.Printf("第%d的x值是：%d\n",i,x)
	}
}

//输出2-100间的素数
func prime_number(){
	//var i int
	//var j int
	var num int
	for i := 1; i<=100;i++{
		num = 0
        for j :=1 ;j<=i;j++{
        	if i%j == 0 {
        		num ++
			}
			}
		if num <=2 {
			fmt.Printf("%d是质数\n",i)
		}
	}
}

func prime_num_b(){
	//
	var i, j int

	for i=2;i<=100;i++{

		for j=2; j<= i/j ; j++{
			if i%j == 0 {
				break
			}
		}
		if (j>(i/j)) {
		    fmt.Printf("%d是素数",i)
		}
	}
}