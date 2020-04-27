package ch16

import (
	"errors"
	"strings"
)

//Pipe-Filter 适用于数据处理和分析
//Filter封装数据处理的功能 松耦合（只和数据格式耦合)
//Pipe用于连接Filter进行数据传递或异步过程中缓冲数据流；进程内同步调用时，用于数据在方法调用间传递

var SplitFilterWrongFormatError = errors.New("Input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter{
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error){
	str, ok := data.(string)//检查数据类型
	if !ok{
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}