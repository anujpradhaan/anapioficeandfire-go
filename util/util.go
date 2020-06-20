package util

import "reflect"

type Root struct {
	Books      string `json:"books"`
	Characters string `json:"characters"`
	Houses     string `json:"houses"`
}

func GetFilterMap(filter interface{}) map[string]string {
	m := make(map[string]string)
	v := reflect.ValueOf(filter)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		value = value.(string)
		if value != "" {
			m[typeOfS.Field(i).Tag.Get("tag")] = value.(string)
		}
	}
	return m
}
