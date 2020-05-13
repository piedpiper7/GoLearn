package ch13

import (
	"errors"
	"time"
)

type ReusableObj struct {

}//预置对象

type ObjPool struct {
	bufChan chan *ReusableObj//用于缓存对象
}

func NewObjPool(numObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numObj)//创建channel
	for i:=0; i<numObj; i++{
		objPool.bufChan <- &ReusableObj{}//将对象放入缓冲池
	}
	return &objPool
}

//获得对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error){//池的方法，定义在池的对象指针上
	select {
	case ret := <-p.bufChan:
		return ret, nil
		case <- time.After(timeout):
			return nil, errors.New("time out")
	}
}

//释放对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}