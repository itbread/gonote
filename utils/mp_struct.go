package utils

import (
	"errors"
	"reflect"
	"strconv"
)

var emptyErr = errors.New("data empty")
var emptyTitleErr = errors.New("title data empty")

//ConvertObject 将一行数据转换
func ConvertObject(row []string, titlesMap map[int]string, chEnMaps map[string]string, obj interface{}) error {
	if len(row) < 1 {
		return emptyErr
	}
	if len(titlesMap) < 1 {
		return emptyTitleErr
	}

	objmap := map[string]string{}
	for index, colCell := range row {
		if ok, colName := convertIndexToItemName(chEnMaps, titlesMap, index); ok && len(colName) > 0 {
			objmap[colName] = colCell
		}
	}
	if len(objmap) > 0 {
		return BuildStruct(obj, objmap)
	}

	return emptyErr

}

//convertIndexToItemName 根据列索引获取列真实字段名称
func convertIndexToItemName(titleTempMap map[string]string, titles map[int]string, index int) (bool, string) {
	match := false
	vl := ""
	if len(titleTempMap) < 1 {
		if key, ok := titles[index]; ok {
			vl = key
		}
	} else {
		if key, ok := titles[index]; ok {
			if title, ok := titleTempMap[key]; ok {
				vl = title
			}
		}
	}
	match = len(vl) > 0
	return match, vl
}

//BuildStruct 将map中对应key的值设置到结构体对应字段上
func BuildStruct(SrcStructPtr interface{}, mp map[string]string) error {
	srcv := reflect.ValueOf(SrcStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	if srct.Kind() != reflect.Ptr || srct.Elem().Kind() == reflect.Ptr {
		return errors.New("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() {
		return errors.New("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	srcfields := DeepFields(srcV.Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := srcV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		vl, ok := mp[v.Name]
		if !ok {
			continue
		}
		switch src.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intNumber, err := strconv.ParseInt(vl, 10, 64)
			if err != nil {
				continue
			}
			dst.SetInt(intNumber)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintNumber, err := strconv.ParseUint(vl, 10, 64)
			if err != nil {
				continue
			}
			dst.SetUint(uintNumber)
		case reflect.Float32, reflect.Float64:
			floatNumber, err := strconv.ParseFloat(vl, 64)
			if err != nil {
				continue
			}
			dst.SetFloat(floatNumber)
		case reflect.Bool:
			tmpOk, err := strconv.ParseBool(vl)
			if err != nil {
				continue
			}
			dst.SetBool(tmpOk)
		case reflect.String:
			dst.SetString(vl)
		}
	}
	return nil
}

//DeepFields 获取 结构体所有的字段
func DeepFields(reflectType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < reflectType.NumField(); i++ {
		v := reflectType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}
