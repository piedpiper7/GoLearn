package ch16

import (
	"errors"
	"strconv"
)

var toIntFilterFormatError = errors.New("Input must be string")

type ToINtFilter struct {

}

func NewToIntFilter() *ToINtFilter{
	return &ToINtFilter{}
}

func (tif *ToINtFilter) Process(data Request)(Response, error){
	parts, ok := data.([]string)
	if !ok {
		return nil, toIntFilterFormatError
	}
	ret := []int{}
	for _,part := range parts{
		s, err := strconv.Atoi(part)
		if err != nil{
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}