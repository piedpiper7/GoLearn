package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

//response
type SignUp struct {
	Succeed bool `json:"succeed"`
	SessionId string `json:"session_id"`
}

//data model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}

type Comment struct {
	Id string
	VideoId string
	Author   string
	Content string
}

type SimpleSession struct {
	UserName string//login_name
	TTL int64
}