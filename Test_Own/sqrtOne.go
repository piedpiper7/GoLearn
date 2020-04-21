package main
import "fmt"
/*
根据牛顿迭代法，举个例子：
例如10，找3*3和4*4
（3+10/3）/2<3
等价于 (i+x/i)/2=i   等价于i-(i*i+x)/2*i=0  等价于(i*i-x)/(2*i) = 0,
当remain趋向于0，说明越接近开方值
*/
func sqrt(x float64,i float64) (float64,float64){
	remain:=(i*i-x)/(2*i);
	i=i-remain
	if(remain>0){
		return sqrt(x,i);
	}else{
		return i,remain
	}
}
func get_sqrt(x float64) float64{
	i,_ :=sqrt(x,x);
	return i;
}
func main(){
	fmt.Println(get_sqrt(4))
	fmt.Println(get_sqrt(3))
}