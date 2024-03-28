package main

import "reflect"

func getFieldName(obj interface{}, field interface{}) string {
	objValue := reflect.ValueOf(obj)
	fieldValue := reflect.ValueOf(field)

	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	objType := objValue.Type()

	for i := 0; i < objType.NumField(); i++ {
		if objValue.Field(i).Pointer() == fieldValue.Pointer() {
			return objType.Field(i).Name
		}
	}
	return ""
}

func getJsonTagName[parent any, child any](p *parent, c *child) string {
	t := reflect.TypeOf(*p)
	s, _ := t.FieldByName(getFieldName(p, c))
	return s.Tag.Get("json")
}
