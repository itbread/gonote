package utils

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
)

//GetDataGrid 获取表格数据
func GetDataGrid(fileName string, sheetName string) ([][]string, error) {
	var data [][]string
	var err error
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	tmpsheetName := sheetName
	if len(tmpsheetName) < 1 {
		tmpsheetName = f.GetSheetName(1)
	}

	return f.GetRows(tmpsheetName)
}

//ReadTmpFromJson json 文件转map
//读取json文件转成map
func ReadTmpFromJson(fileName string) map[string]string {
	mp := make(map[string]string)
	if contents, err := ioutil.ReadFile(fileName); err == nil {
		er := json.Unmarshal(contents, &mp)
		if er != nil {
			fmt.Println("Unmarshal err:", err)
		}
	} else {
		fmt.Println("read file err:", err)
	}
	fmt.Println("ReadTmpFromJson mp============", mp)
	return mp

}

// SliceToMap slice 转 map
//读取标题栏 列索引与字段对应关系
func SliceToMap(titles []string) map[int]string {
	mp := make(map[int]string)
	for index, title := range titles {
		if len(title) > 0 {
			mp[index] = title
		}
	}
	return mp
}
