package ch16

type Request interface {

}

type Response interface {

}

//接口的方法 pipe-filter structure
type Filter interface {
	Process(data Request)(Response, error)
}