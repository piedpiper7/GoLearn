package ch11

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConCurrentMap(t *testing.T){
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 100)
	t.Log(m.Get(cm.StrKey("key")))
}
