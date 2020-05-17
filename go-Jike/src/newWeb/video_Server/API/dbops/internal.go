package dbops

import (
	"database/sql"
	"log"
	"newWeb/video_Server/API/defs"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error{
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions(session_id, TTL, login_name) VALUE (?, ?, ?)")
	if err != nil{
		return err
	}

	_, err = stmtIns.Exec(&sid, &ttlStr, &uname)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error){
	session := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT  TTL, login_name FROM sessions WHERE session_id = ?")
	if err != nil{
		return nil, err
	}
	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil{
		session.TTL = res
		session.UserName = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return session, nil
}

func RetrieveAllSession() (*sync.Map, error) {//返回的是全部session的k-v
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil{
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil{
		log.Printf("%s", err)
		return nil, err
	}
	for rows.Next(){
		var id, ttlStr, uname string
		if err := rows.Scan(&id, &ttlStr, &uname); err != nil{
			log.Printf("retrieve session error:%s", err)
			break
		}
		if ttl, err := strconv.ParseInt(ttlStr, 10, 64); err == nil{
			session := &defs.SimpleSession{UserName: uname, TTL: ttl}
			m.Store(id, session)
			log.Printf("session id %s, ttl %s", id, session.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error{
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	if err != nil{
		log.Printf("%s", err)
		return err
	}
	if _, err := stmtDel.Exec(sid); err != nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}