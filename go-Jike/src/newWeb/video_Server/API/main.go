package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	vaildateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}

func main(){
	r := RegisterHandlers()
	middleHandler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", middleHandler)
}

//ListenAndServe用于注册，会注册r的handlers，每个handlers都是用不同的groutine(CreateUser/Login)
//groutine是一种轻量级的类线程的协程
