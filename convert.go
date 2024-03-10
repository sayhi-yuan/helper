package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// 各种符合类型的转换

// ConvertStructToMap 结构体转成map, 只支持一层结构
func ConvertStructToMap(s any) (result map[string]any, err error) {
	result = map[string]any{}

	v := reflect.ValueOf(s)
	// 判断是否为结构体
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		err = fmt.Errorf("传入的参数必须是struct或者*struct")
		return
	}

	// 转换的key为结构体的tag->json
	for i := 0; i < v.NumField(); i++ {
		// 获取tag
		tag := v.Type().Field(i).Tag.Get("json")
		if tag != "" {
			result[tag] = v.Field(i).Interface()
		}
	}

	return
}

// ConvertMapToStruct map转成结构体, 只支持一层结构
// m 为 map 类型
// s 为 *struct 类型
func ConvertMapToStruct(m any, s any) (err error) {
	// 检测类型
	mV := reflect.ValueOf(m)
	if mV.Kind() == reflect.Ptr {
		mV = mV.Elem()
	}
	if mV.Kind() != reflect.Map {
		err = fmt.Errorf("参数m必须是map类型")
		return
	}

	sV := reflect.ValueOf(s)
	if sV.Kind() != reflect.Ptr || sV.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("参数s必须是*struct类型")
		return
	}

	b, err := json.Marshal(m)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, s)

	return
}
