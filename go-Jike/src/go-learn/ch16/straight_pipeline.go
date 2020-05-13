package ch16

type StraightPipeLine struct {
	Name string
	Filters *[]Filter
}

func NewStraightPipeLine(name string, filters...Filter) *StraightPipeLine{
	return &StraightPipeLine{
		name,
		&filters,
	}
}

func (f *StraightPipeLine) Process(data Request)(Response,error){
	var ret interface{}
	var err error
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)//依次调用
		if err != nil {
			return ret, err
		}
		data = ret//上一个的结果作为下一个的输入
	}
	return data, err
}

