package main

import (
	"fmt"
	"log"
	"reflect"
)

//定义控制器函数Map类型，便于后续快捷使用
type ControllerMapsType map[string]reflect.Value

//声明控制器函数Map类型变量
var ControllerMaps ControllerMapsType

//定义路由器结构类型
type User struct {
	username string
	password int64
	email    string
}

//获取结构体中字段的名称
func GetFieldNames(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}

//获取结构体中字段的内容
//获取结构体中字段的内容
func GetStructFields(structName interface{}) []reflect.Value {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}
	fieldNum := t.NumField()
	structValue := reflect.ValueOf(structName)
	var result []reflect.Value
	for i := 0; i < fieldNum; i++ {
		rst := structValue.FieldByName(t.Field(i).Name)
		result = append(result, rst)
	}
	fmt.Printf("%v is int64 type\n", result)

	return result
}

func main() {
	a := User{}
	a.email = "1191500820"
	a.username = "yancy"
	a.password = 12434
	GetStructFields(a)

	//fmt.Println(strings.Join(str,","))

}
