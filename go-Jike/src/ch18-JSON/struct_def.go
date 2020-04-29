package ch18_JSON

//easyjson -all filename

type BasicInfo struct {
	Name string `json:"name"`//利用json映射
	Age int `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo JobInfo `json:"job_info"`
}