package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name",Login)
	return router
}

func main(){
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

//ListenAndServe用于注册，会注册r的handlers，每个handlers都是用不同的groutine(CreateUser/Login)
//groutine是一种轻量级的类线程的协程
