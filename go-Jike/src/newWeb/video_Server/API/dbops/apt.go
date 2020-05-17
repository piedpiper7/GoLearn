package dbops

import (
	"database/sql"
	"log"
	"newWeb/video_Server/API/defs"
	"newWeb/video_Server/API/utils"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error){
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("UserGet Errpr:%s", err)
		return "", nil
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}

func UserDelete(loginName string, pwd string) error{
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil{
		log.Printf("UserDelete Error:%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	defer stmtDel.Close()

	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error){
	vid, err := utils.NewUUID()
	if err != nil{
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO video_info (id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)")
	if err != nil{
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil{
		return nil, err
	}
	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error){
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	var(
		aid int
		name string
		ctime string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows{
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil{
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil{
		return err
	}
	defer stmtDel.Close()
	return nil
}


func AddNewComment(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare("INSERT INTO comments(id, video_id, author_id, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil{
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error){
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content
                                           FROM comments INNER JOIN users ON comments.author_id = users.id
                                           WHERE comments.video_id = ? and comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil{
		return res, err
	}
	for rows.Next(){
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil{//利用迭代器把每一行内容取出赋值给变量
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}