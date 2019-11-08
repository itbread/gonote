package services

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
)

//获取sql like 查询条件
func BuildLikeCondition(db *gorm.DB, key string, value string, tableName ...string) *gorm.DB {
	if len(key) < 1 {
		return db
	}
	fullKey := ""
	if len(tableName) > 0 {
		fullKey = fmt.Sprintf("%v.%v", tableName[0], key)
	} else {
		fullKey = key
	}
	if len(value) > 1 {
		query := fmt.Sprintf("%v like ?", fullKey)
		args := fmt.Sprintf("%v%v%v", "%", value, "%")
		db = db.Where(query, args)
	}
	return db
}

//获取sql Equal 查询条件
// 注意 value 只支持 int string bool类型
func BuildEqualCondition(db *gorm.DB, key string, value interface{}, tableName ...string) *gorm.DB {
	if len(key) < 1 {
		return db
	}
	query := ""
	fullKey := key
	if len(tableName) > 0 {
		fullKey = fmt.Sprintf("%v.%v", tableName[0], key)
	}

	switch value.(type) {
	case float32, float64, int, int32, int64, uint, uint32, uint64:
		query = fmt.Sprintf("%v = ?", fullKey)
		db = db.Where(query, value)
	case bool:
		query = fmt.Sprintf("%v = ?", fullKey)
		db = db.Where(query, value)
	case string:
		vl := value.(string)
		if len(vl) > 1 {
			query = fmt.Sprintf("%v = ?", fullKey)
			db = db.Where(query, value)
		}
	}

	fmt.Printf("======fullKey=%v =query: %v args: %v \n", fullKey, query, value)
	return db
}

//生成 sql in 条件 value 是 Slice
func BuildInCondition(db *gorm.DB, key string, value interface{}, tableName ...string) *gorm.DB {
	if len(key) < 1 {
		return db
	}
	fullKey := ""
	if len(tableName) > 0 {
		fullKey = fmt.Sprintf("%v.%v", tableName[0], key)
	} else {
		fullKey = key
	}
	srcv := reflect.ValueOf(value)
	//判断是否为 Slice
	if srcv.Kind() != reflect.Slice {
		return db
	}
	//判断 Slice 是否为空
	if srcv.Len() < 1 {
		return db
	}

	query := fmt.Sprintf("%v in ?", fullKey)
	db = db.Where(query, value)

	return db
}

//获取sql Between 查询条件
func BuildBetweenCondition(db *gorm.DB, key string, min interface{}, max interface{}, tableName ...string) *gorm.DB {
	if len(key) < 1 {
		return db
	}

	fullKey := key
	if len(tableName) > 0 {
		fullKey = fmt.Sprintf("%v.%v", tableName[0], key)
	}

	fmt.Printf("======fullKey=%v =min: %v max: %v \n", fullKey, min, max)
	if min != nil && max != nil {
		query := fmt.Sprintf("%v between ? and ?", fullKey, min, max)
		db = db.Where(query, min, max)
	}
	return db
}
