package main

import (
	"net/http"
	"newWeb/video_Server/API/defs"
	"newWeb/video_Server/API/session"
)

//authenticate

//自定义header
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func vaildateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return  false
	}
	uname, ok := session.IsSessionExpired(sid); if ok{
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func validateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0{
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}