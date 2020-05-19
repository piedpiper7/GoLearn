package streamserver

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *LimiterBucket
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler{
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandler() *httprouter.Router{
	r := httprouter.New()
	r.GET("/videos/:vid-id", streamHandler)
	r.POST("upload/:vid-id", uploadHandler)

	return r
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if !m.l.GetConn(){
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many request")
		return
	}
	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main(){
	r := RegisterHandler()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}