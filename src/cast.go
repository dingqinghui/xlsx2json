/**
 * @Author: dingQingHui
 * @Description:
 * @File: cast
 * @Version: 1.0.0
 * @Date: 2024/10/28 14:41
 */

package src

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/spf13/cast"
	"log"
	"strings"
)

func Cast(sheet string, row, column int, meta *fieldMeta, v string) interface{} {
	if v == "" {
		v = meta.DataDefault
	}

	if meta.IsSlice() {
		return sliceCast(sheet, row, column, v, meta)
	}

	// 枚举转换
	if enumMeta, ok := GetXlsxStructHub().GetEnumMeta(meta.DataType, v); ok {
		if enumMeta == nil {
			log.Printf("未知的枚举类型 sheet:%v row:%d column:%d enum:%v value:%v \n", sheet, row, column, meta.DataType, v)
			return 0
		}
		return cast.ToInt32(enumMeta.DataDefault)
	}

	switch meta.DataType {
	case "int":
		return cast.ToInt32(v)
	case "int64":
		return cast.ToInt64(v)
	case "float":
		return cast.ToFloat32(v)
	case "float64":
		return cast.ToFloat64(v)
	case "string":
		return v
	case "bool":
		return cast.ToBool(v)
	default:
		return structCast(sheet, row, column, v, meta.DataType)
	}
}

func sliceCast(sheet string, row, column int, data string, meta *fieldMeta) *gabs.Container {
	if data == "" {
		return nil
	}
	jsonObj := gabs.New()
	array, _ := jsonObj.Array()
	values := strings.Split(data, meta.DataSlicing)
	copyGlobalCellTable := meta.Copy()
	copyGlobalCellTable.DataSlicing = ""
	for _, value := range values {
		v := Cast(sheet, row, column, copyGlobalCellTable, value)
		if err := array.ArrayAppend(v); err != nil {
			log.Fatal(data, err)
			return nil
		}
	}
	return jsonObj
}

// String2JsonStruct
// @Description: 将字符串转为Json结构体并序列化
// @param value
// @param structName
// @return string
func structCast(sheet string, row, column int, data string, structName string) *gabs.Container {
	if data == "" {
		return nil
	}
	a := GetXlsxStructHub()
	_ = a
	jsonStruct := gabs.New()
	fieldValues := strings.Split(data, ",")
	for _, fieldValue := range fieldValues {
		infos := strings.Split(fieldValue, ":")
		if len(infos) < 2 {
			log.Fatalf("structCast sheet:%v row:%d column:%d name:%v value:%v\n", sheet, row, column, structName, fieldValue)
		}
		fieldName := infos[0]
		fieldData := infos[1]
		s := GetXlsxStructHub().Struct
		_ = s
		_meta := GetXlsxStructHub().GetMeta(structName, fieldName)
		v := Cast(sheet, row, column, _meta, fieldData)
		if _, err := jsonStruct.SetP(v, _meta.ObjName); err != nil {
			log.Fatal(structName, data, err)
		}
	}
	return jsonStruct
}
