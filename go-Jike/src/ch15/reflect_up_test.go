package ch15

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeID string
	Name string `format:"normal"` //struct tag (key-value) 和json对应
	Age int
}

type Custmer struct {
	CookieID string
	Name string
	Age int
}

func (e *Employee) UpdateAge(newVal int){
	e.Age = newVal
}

func TestInvokeByName(t *testing.T){
	e := &Employee{"1", "Mary", 18}
	t.Logf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	//通过Name利用反射返回value,value只返回一个值 但是Type返回两个值，需要判断
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok{
		t.Error("Error to get field name.")
	} else {
		t.Log("Tag:format",nameField.Tag.Get("format"))}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(11)})
	t.Log("Update Age ", e)
}

func FillBySettings(st interface{}, settings map[string] interface{} ) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr{
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct{//elem()获取指针指向的值
			return errors.New("the first param should be pointer to the struct type")
		}
	}
	if settings == nil{
		return errors.New("the map should have value")
	}

	var (
		field reflect.StructField
		ok bool
		)

	for k,v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v){
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T){
	settings := map [string] interface{}{"Name":"Mike", "Age":17}
	e := Employee{}
	if err := FillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Custmer)
	if err := FillBySettings(c, settings); err != nil{
		t.Fatal(err)
	}
	t.Log(*c)
}