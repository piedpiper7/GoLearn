package Series

func FibonacciGrt(n int) [] int {
	var FibList = [] int {1, 1}
	for i:=2; i<=n; i++ {
		FibList = append(FibList, FibList[i-1] + FibList[i-2])
	}
	return  FibList
}

func Square(n int) int {
	return n * n
}