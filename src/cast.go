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

func Cast(meta *fieldMeta, v string) interface{} {
	if v == "" {
		v = meta.DataDefault
	}

	// 枚举转换为int32
	if enumMap, ok := GetXlsxStructHub().Enum[meta.DataType]; ok {
		enumMeta := enumMap[v]
		return cast.ToInt32(enumMeta.DataDefault)
	}

	if meta.IsSlice() {
		return sliceCast(v, meta)
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
	default:
		return structCast(v, meta.DataType)
	}
}

func sliceCast(data string, meta *fieldMeta) *gabs.Container {
	if data == "" {
		return nil
	}
	jsonObj := gabs.New()
	array, _ := jsonObj.Array()
	values := strings.Split(data, meta.DataSlicing)
	copyGlobalCellTable := meta.Copy()
	copyGlobalCellTable.DataSlicing = ""
	for _, value := range values {
		v := Cast(copyGlobalCellTable, value)
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
func structCast(data string, structName string) *gabs.Container {
	if data == "" {
		return nil
	}
	jsonStruct := gabs.New()
	fieldValues := strings.Split(data, ",")
	for _, fieldValue := range fieldValues {
		infos := strings.Split(fieldValue, ":")
		if len(infos) < 2 {
			log.Fatalf("structCast name:%v value:%v\n", structName, fieldValue)
		}
		fieldName := infos[0]
		fieldData := infos[1]
		_meta := GetXlsxStructHub().GetMeta(structName, fieldName)
		v := Cast(_meta, fieldData)
		if _, err := jsonStruct.SetP(v, _meta.ObjName); err != nil {
			log.Fatal(structName, data, err)
		}
	}
	return jsonStruct
}
