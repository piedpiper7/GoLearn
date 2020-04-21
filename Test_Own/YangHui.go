package main

import "fmt"
/*
思路是：创建两个数组，数组B作为初始化数组，初始化为1，将B数组赋值给C，再清空B数组
通过遍历数组C，将B按照杨辉三角的格式append
 */
/*
 */
func main(){
	YangHUiTriangle()
}

func YangHUiTriangle() {
	var arrayB [] int
	arrayB = append(arrayB, 1)
	for x := 1; x < 20; x++ {
		//var arrayC [] int
		/*for i := 0; i < len(arrayB); i++ {
			arrayC = append(arrayC, arrayB[i])
		}*/

	var arrayC = make([]int, len(arrayB),cap(arrayB))
	copy(arrayC, arrayB)
	//为什么直接赋值不可以  用切片复制
	//arrayC = arrayB
	//清空数组arrayB的元素
	for i := 0; i < len(arrayB); i++ {
		arrayB = append(arrayB[:i])
	}

		for i := 0; i < x; i++ {
			var num1 int
			if i == 0 {
				num1 = arrayC[i]
			} else if i == len(arrayC) {
				num1 = arrayC[i-1]
			} else {
				num1 = arrayC[i-1] + arrayC[i]
			}
			arrayB = append(arrayB, num1)
			fmt.Printf("[%d]", arrayB[i])

		}
		fmt.Print("\n")
	}
}