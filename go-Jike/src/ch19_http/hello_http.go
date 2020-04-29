package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello World!")
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request){
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write([]byte(timeStr))//写回request
	})
	http.ListenAndServe(":8080", nil)
}

//以/结尾的url可以匹配它的任何子路径，采用最长匹配原则
//末尾不是/则说明是叶子节点