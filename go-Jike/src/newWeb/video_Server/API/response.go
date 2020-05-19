package main

import (
	"encoding/json"
	"io"
	"net/http"
	"newWeb/video_Server/API/defs"
)

func sendErrorResponse(w http.ResponseWriter, errRes defs.ErrorResponse) {
	w.WriteHeader(errRes.HttpSC)
	resStr, _ := json.Marshal(&errRes.Error)//序列化为json
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int){
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
