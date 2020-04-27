package ch16

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPipeLine(t *testing.T){
	split := NewSplitFilter(",")
	convert := NewToIntFilter()
	sum := NewSumFilter()
	sp := NewStraightPipeLine("pl", split, convert, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil{
		t.Fatal(err)
	}
	t.Log(ret)
	assert.Equal(t,ret,6)
}
