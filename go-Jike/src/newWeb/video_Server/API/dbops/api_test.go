package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables(){
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M){
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T){
	t.Run("ADD", TestAddUser)
	t.Run("GET", TestGetUser)
	t.Run("DELETE", TestDeleteUser)
	t.Run("REGET", TestRegetUser)
}

func TestAddUser(t *testing.T){
	err := AddUserCredential("sunny", "123")
	if err != nil{
		t.Errorf("Add user error:%v", err)
	}
}

func TestGetUser(t *testing.T){
	pwd, err := GetUserCredential("sunny")
	if err != nil || pwd != "123" {
		t.Errorf("Get user error:%v", err)
	}
}

func TestDeleteUser(t *testing.T){
	err := UserDelete("sunny", "123")
	if err != nil {
		t.Errorf("Delete User error:%v", err)
	}
}

func TestRegetUser(t *testing.T){
	pwd, err := GetUserCredential("sunny")
	if err != nil{
		t.Errorf("Reget user error :%v", err)
	}
	if pwd != ""{
		t.Errorf("Delete user error:%v", err)
	}
}

func TestVideoWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser", TestAddUser)
	t.Run("ADDVideo", testAddVideo)
	t.Run("GETVideo", testGetVideo)
	t.Run("DELETEVideo", testDeleteVideo)
	t.Run("REGETVideo", testRegetVideo)
}

var tempVid string

func testAddVideo(t *testing.T){
	vi, err := AddNewVideo(1, "my-video")
	if err != nil{
		t.Errorf("Add new video error: %v", err)
	}
	tempVid = vi.Id
}

func testGetVideo(t *testing.T){
	_, err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Get Video info error: %v", err)
	}
}

func testDeleteVideo(t *testing.T){
	err := DeleteVideoInfo(tempVid)
	if err != nil{
		t.Errorf("Delete video error: %v", err)
	}
}

func testRegetVideo(t *testing.T){
	vi, err := GetVideoInfo(tempVid)
	if err != nil || vi != nil{
		t.Errorf("Reget video error: %v", err)
	}
}

func TestCommentWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser", TestAddUser)
	t.Run("AddComment", testAddComment)
	t.Run("ListComment", testListComment)
}

func testAddComment(t *testing.T){
	err := AddNewComment("123", 1, "this is very wonderful")
	if err != nil{
		t.Errorf("Add new comment error: %v", err)
	}
}

func testListComment(t *testing.T){
	vid := "123"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	cm, err := ListComments(vid, from, to)
	if err != nil{
		t.Errorf("List comment error: %v", err)
	}
	for i, ele := range cm{
		fmt.Printf("Comment %d, %v \n", i, ele)
	}
}