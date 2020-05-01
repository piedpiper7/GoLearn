package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Welcome\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params){//httprouter声明的参数在byNmae获得
	fmt.Fprintf(w, "Hello,%s!\n", ps.ByName("name"))
}

func main(){
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
