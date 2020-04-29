package ch18_JSON

import (
	"encoding/json"
	"fmt"
	"testing"
)

var JsonStr =`{
"basic_info":{
"name":"Mike",
"age":17},
"job_info":{
"skills":["Java","Go","C"]}
}`

func TestEmbeddedJson(t *testing.T){//嵌入式Json
	e := new(Employee)
	err := json.Unmarshal([]byte(JsonStr),e)//将json填充到对象中
	if err != nil{
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil{
		fmt.Println(string(v))
	} else{
		t.Error(err)
	}
}

func TestEasyJson(t *testing.T){
	e := Employee{}
	e.UnmarshalJSON([]byte(JsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err == nil{
		fmt.Println(string(v))
	} else {
	t.Error(err)}
}

func BenchmarkEmbeddedJson(b *testing.B){
	b.ResetTimer()
	e := new(Employee)//new创建返回的是一个指针，分配的是类型
	for i:=0; i<b.N; i++{
		err := json.Unmarshal([]byte(JsonStr),e)
		if err != nil{
			b.Error(err)
		}
		if _,err := json.Marshal(e); err != nil{
			b.Error(err)
		}
	}
}

func BenchmarkEasyJson(b *testing.B){
	b.ResetTimer()
	e := Employee{}//返回的是struct类型的值
	for i:=0; i<b.N; i++{
		err := e.UnmarshalJSON([]byte(JsonStr))
		if err != nil{
			b.Error(err)
		}
		if _, err = e.MarshalJSON(); err != nil{
			b.Error(err)
		}
	}

}

//terminal 使用 go test -bench=.