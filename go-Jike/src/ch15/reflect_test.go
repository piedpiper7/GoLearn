package ch15

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}){
	n := reflect.TypeOf(v)
	switch n.Kind() {//kind 是一个const 里面包含全部类型
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", n)
	}
}

func TestTypeCheck(t *testing.T){
	var f = float32(10)
	CheckType(f)
	var g float64 = 20
	CheckType(&g)
}

func TestTypeAndValue(t *testing.T){
	var f float64 = 7
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}