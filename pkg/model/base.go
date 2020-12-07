package model

import (
	"fmt"
	"go-api/api/pkg/database"

	"gorm.io/gorm"
)

// 分页数据
type PageInfo struct {
	Page     int64
	PageSize int64
}

// 分页返回数
type PageList struct {
	CurrentPage int64
	FirstPage   int64
	LastPage    int64
	PageSize    int64
	Total       int64
	Data        interface{}
}

// 查询条件
type Condition struct {
	Key   string
	Value string
	Way   string
}

/**
设置condition 查询条件数据
*/
func SetCondition(params []Condition, key string, val string, where ...string) []Condition {
	condition := Condition{Key: key, Value: val}
	if where != nil {
		condition.Way = where[0]
	} else {
		condition.Way = "="
	}
	params = append(params, condition)
	return params
}

func orm(condition []Condition) *gorm.DB {
	query, values := formatQuery(condition)
	rom := database.DB
	return getConditionOrm(rom, query, values)
}

func formatQuery(condition []Condition) (string, []interface{}) {
	query := ""
	n := len(condition)
	values := make([]interface{}, n)
	for ix, value := range condition {
		if query == "" {
			query = fmt.Sprintf("%s %s ?", value.Key, value.Way)
		} else {
			query = fmt.Sprintf("%s and %s %s ?", query, value.Key, value.Way)
		}
		values[ix] = value.Value
	}
	return query, values
}

func getConditionOrm(rom *gorm.DB, query string, values []interface{}) *gorm.DB {
	switch len(values) {
	case 1:
		rom = rom.Where(query, values[0])
	case 2:
		rom = rom.Where(query, values[0], values[1])
	case 3:
		rom = rom.Where(query, values[0], values[1], values[2])
	case 4:
		rom = rom.Where(query, values[0], values[1], values[2], values[3])
	case 5:
		rom = rom.Where(query, values[0], values[1], values[2], values[3], values[4])
	case 6:
		rom = rom.Where(query, values[0], values[1], values[2], values[3], values[4], values[5])
	case 7:
		rom = rom.Where(query, values[0], values[1], values[2], values[3], values[4], values[5], values[6])
	case 8:
		rom = rom.Where(query, values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7])
	case 9:
		rom = rom.Where(query, values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7], values[8])
	case 10:
		rom = rom.Where(query, values[0], values[1], values[2], values[3], values[4], values[5], values[6], values[7], values[8], values[9])
	}
	return rom
}
