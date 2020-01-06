package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Name string
	Age  int
	Num  int64
	Id   uint64
	Cource
}
type Cource struct {
	CourceName string
	CourceCode int
}

func getTag(tag string) string {
	const tag_const = "json:\"\""
	return tag
}

func debug(tag, obj interface{}) {
	if bts, err := json.Marshal(obj); err == nil {
		fmt.Println("############------", tag, "  ", string(bts))
	} else {
		fmt.Println("############----err=", tag, "  ", err)
	}

}

func BuildObject(i interface{}, mp map[string]string) {
	//获取指针指向的真正的数值Value
	valueOfI := reflect.ValueOf(i).Elem()
	//获取对应的Type这个是用来获取属性方法的
	typeOfI := valueOfI.Type()
	//判断是否是struct
	if typeOfI.Kind() != reflect.Struct {
		fmt.Println("except struct")
		return
	}
	//获取属性的数量
	numField := typeOfI.NumField()
	//遍历属性，找到特定的属性进行操作
	for i := 0; i < numField; i++ {
		//获得属性的StructField，次方法不同于Value中的Filed（这个返回的是Field）
		field := typeOfI.Field(i)
		debug("item ", field)

		//获取属性 名称
		fieldName := field.Name
		//fildTag := field.Tag
		//anonymous := field.Anonymous
		fmt.Println("fildName ======", fieldName)
		//fmt.Println("fildTag  ======", fildTag)
		//fmt.Println("anonymous======", anonymous)
		fmt.Println("valueOfI.Field(i).Kind()======", valueOfI.Field(i).Kind())
		vl, ok := mp[fieldName]
		fmt.Println("vl======", vl, "  ok==", ok)
		if !ok {
			continue
		}
		kind := valueOfI.Field(i).Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			int_numbe, err := strconv.ParseInt(vl, 10, 64);
			fmt.Println("int_numbe======", int_numbe, "  err==", err)
			if err != nil {
				continue
			}
			valueOfI.Field(i).SetInt(int_numbe)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uint_numbe, err := strconv.ParseUint(vl, 10, 64);
			fmt.Println("unit_numbe======", uint_numbe, "  err==", err)
			if err != nil {
				continue
			}
			valueOfI.Field(i).SetUint(uint_numbe)
		case reflect.Float32, reflect.Float64:
			float_numbe, err := strconv.ParseFloat(vl, 64);
			fmt.Println("unit_numbe======", float_numbe, "  err==", err)
			if err != nil {
				continue
			}
			valueOfI.Field(i).SetFloat(float_numbe)
		//case reflect.Uint:
		//case reflect.Uint8:
		//case reflect.Uint16:
		//case reflect.Uint32:
		//case reflect.Uint64:
		case reflect.Bool:
			//valueOfI.Field(i).SetBool(false)
			if tmp_ok, err := strconv.ParseBool(vl); err != nil {
				valueOfI.Field(i).SetBool(tmp_ok)
			}
		case reflect.String:
			fmt.Println("string**********", valueOfI.Field(i).Kind())
			valueOfI.Field(i).SetString(vl)
		case reflect.Struct:

		}

		////找到名为Name的属性进行修改值
		//if fieldName == "Name" {
		//	//改变他的值为jack
		//	valueOfI.Field(i).SetString(vl)
		//
		//	fmt.Println("type======", valueOfI.Field(i).Type())
		//}
	}
}

func main() {
	stu := Student{Name: "susan", Age: 58}
	mp := make(map[string]string)
	BuildObject(&stu, mp)
	fmt.Println("stu====", stu)
	var tmp Student
	mp["Name"] = "itbread"
	mp["Age"] = "30"
	mp["Num"] = "1000"
	mp["Id"] = "123456"
	BuildObject(&tmp, mp)
	fmt.Println("tmp===", tmp)
}
