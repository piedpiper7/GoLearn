package ch16

import "errors"

var SumFilterFormatError = errors.New("Input must be int")

type SumFilter struct {

}

func NewSumFilter() *SumFilter{
	return &SumFilter{}
}

func (su *SumFilter) Process(data Request)(Response, error){
	elems, ok := data.([]int)
	if !ok{
		return nil, SumFilterFormatError
	}
	ret := 0
	for _,elem := range elems{
		ret += elem
	}
	return ret, nil
}