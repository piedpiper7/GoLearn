package session

import (
	"sync"
	"time"
	"newWeb/video_Server/API/dbops"
	"newWeb/video_Server/API/defs"
	"newWeb/video_Server/API/utils"
)

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func nowInMIlli() int64{
	return time.Now().UnixNano()/1000000
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDb(){
	r, err := dbops.RetrieveAllSession()
	if err != nil{
		return
	}

	r.Range(func(k, v interface{}) bool{ //将db得到的session写入sessionMap
		session := v.(*defs.SimpleSession)
		sessionMap.Store(k, session)
		return true
	})
}

func GenerateNewSessionId(uname string) string{
	id, _ := utils.NewUUID()
	ct := nowInMIlli()
	ttl := ct + 30*60*1000 //server session valid time 30 min

	session := &defs.SimpleSession{UserName: uname, TTL: ttl}
	sessionMap.Store(id, session)
	dbops.InsertSession(id, ttl, uname)
	return id
}

func IsSessionExpired(sid string) (string, bool){
	session, ok := sessionMap.Load(sid)
	if ok{
		ct := nowInMIlli()
		if session.(*defs.SimpleSession).TTL < ct{
			deleteExpiredSession(sid)
			return "", true
		}
		return session.(*defs.SimpleSession).UserName, false
	}
	return "", true
}