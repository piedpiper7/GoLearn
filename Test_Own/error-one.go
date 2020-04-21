package main

import "fmt"

type DivideError struct{
	dividee int
	divider int
}

func (de *DivideError) Error() string{
	strFormat :=`
	Cannot proceed, the divider is zero.
	dividee:%d
	divider:0
`

	return fmt.Sprintf(strFormat,de.dividee)
}

//定义运算函数
func Divide(varDividee int, varDivider int) (result int, errormsg string){
	if varDivider == 0 {
		dData := DivideError{
			dividee:varDividee,
			divider:varDivider,
		}
		errormsg = dData.Error()
		return
	} else {
		return varDividee/varDivider,""
	}
}

func main(){
	result,errormsg:=Divide(100,10)
	fmt.Println("result is ",result, ",msg is ",errormsg)
	result1,errormsg1:=Divide(100,0)
	fmt.Println("result is ",result1, ",msg is ",errormsg1)
	// 正常情况
	/*if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}*/

}