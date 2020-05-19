package streamserver

import "log"

//流控制 视频需要上传和下载..以流的方式

type LimiterBucket struct {
	currentConn int
	bucket chan int
}

func NewConnLimiter(chanel int) *LimiterBucket{
	return &LimiterBucket{
		currentConn: chanel,
		bucket: make(chan int, chanel),
	}
}

func (lb *LimiterBucket) GetConn() bool{
	if len(lb.bucket) >= lb.currentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	lb.bucket <- 1
	return  true
}

func (lb *LimiterBucket) ReleaseConn() {
	c := <- lb.bucket
	log.Printf("New connect coming %d", c)
}