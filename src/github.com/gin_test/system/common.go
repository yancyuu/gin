package system

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

//获取结构体中字段的名称
func GetStructFieldNames(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		Logger().WithFields(logrus.Fields{
			"name": "GetStructFieldNames",
		}).Info("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	Logger().WithFields(logrus.Fields{
		"name": "GetStructFieldNames",
	}).Info("result %v\n", result)
	return result
}

//获取结构体中字段的内容
func GetStructFields(structName interface{}) []interface{} {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		Logger().WithFields(logrus.Fields{
			"name": "yancy",
		}).Info("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	structValue := reflect.ValueOf(structName)
	var result []interface{}
	for i := 0; i < fieldNum; i++ {
		rst := structValue.FieldByName(t.Field(i).Name)
		result = append(result, rst)
	}
	Logger().WithFields(logrus.Fields{
		"name": "GetStructFields",
	}).Info("result %v\n", result)
	return result
}
