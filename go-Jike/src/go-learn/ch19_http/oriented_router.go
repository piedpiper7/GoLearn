package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//resource oriented architecture 面向资源的架构

type Employee struct {
	ID string `json:"id"`
	Nmae string `json:"nmae"`
	Age int `json:"age"`
}

var employeeDB map[string] *Employee

func init(){//模拟数据库
	employeeDB = map[string] *Employee{}
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 17}
	employeeDB["Abby"] = &Employee{"e-2","Abby", 18}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	qName := ps.ByName("name")
	var(
		ok bool
		info *Employee
		infoJson []byte
		err error
	)
	if info, ok = employeeDB[qName]; !ok{
		w.Write([]byte("{\"error\":\"Not Found\"}"))
		return
	}
	if infoJson, err = json.Marshal(info); err != nil{
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))
	}
	w.Write(infoJson)
}

func main(){
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/employee/:name", GetEmployeeByName)
	log.Fatal(http.ListenAndServe(":8088", router))
}