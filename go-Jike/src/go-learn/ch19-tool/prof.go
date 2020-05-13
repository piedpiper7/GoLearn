package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	row = 10000
	col = 10000
	)

func fileMatrix(m *[row][col]int){
	s := rand.New((rand.NewSource(time.Now().UnixNano())))

	for i:=0; i<row; i++{
		for j:=0; j<col; j++{
			m[i][i] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int){
	for i:=0; i<row; i++{
		tmp := 0
		for j:=0; j<col; j++{
			tmp += m[i][j]
		}
	}
}

func main(){
	//创建输出文件
	f, err := os.Create("cpu.prof")
	if err != nil{
		log.Fatal("Could not create CPU profile:", err)
	}
	//获取系统信息
	if err := pprof.StartCPUProfile(f); err != nil{//监控CPU
		log.Fatal("Could not start CPU r=profile:", err)
	}
	defer pprof.StopCPUProfile()

	x := [row][col]int{} //初始化二维数组
	fileMatrix(&x)
	calculate(&x)

	f1, err := os.Create("mem.prof")
	if err != nil{
		log.Fatal("Could not create memory profile:", err)
	}
	runtime.GC()  //获取最新数据信息
	if err := pprof.WriteHeapProfile(f1); err != nil{//写入内存信息 API dump
		log.Fatal("Could not write memory profile:", err)
	}
	f1.Close()

	f2, err := os.Create("goroutine.prof")
	if err != nil{
		log.Fatal("Could not create goroutine profile")
	}

	if gprof := pprof.Lookup("goroutine"); gprof == nil{//profile 不同的tag
		log.Fatal("Could not write goroutine profile:")
	} else {
		gprof.WriteTo(f2, 0)
	}
	f2.Close()
}
 //gp build prof.go
 //go run prof.go
 //go tool pprof prof cpu.prof
 //top  (top CPU的情况)  flat 函数本身的执行时间， cum 调用了其他函数加起来的执行时间
 //list+函数名  显示详细执行情况
//svg graphviz cpu图  8min
//go-torch