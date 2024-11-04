/**
 * @Author: dingQingHui
 * @Description:
 * @File: gen_proto
 * @Version: 1.0.0
 * @Date: 2024/10/28 18:27
 */

package src

import (
	"cmp"
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
	"log"
	"slices"
	"strings"
)

var (
	// proto文件头
	protoHeader = `syntax = "proto3";
option go_package = ".;staticdata";
package Game;

`
)

var (
	xlsxType2ProtoTypeMap = map[string]string{
		"int": "int32",
	}
)

func GenProtoFileCode() {
	var data []string
	data = append(data, protoHeader)

	data = append(data, genProtoFile(GetXlsxStructHub().TableHeader)...)
	data = append(data, genProtoFile(GetXlsxStructHub().Struct)...)

	if err := WriteFile(tmpOutProtoFile, []byte(strings.Join(data, "\n"))); err != nil {
		log.Fatal("GenProtoFile:", tmpOutProtoFile, err)
		return
	}

	log.Println("proto文件生成成功", tmpOutProtoFile)
}

type protoField struct {
	valueType string // 字段类型
	fieldName string // 字段名字
	mark      string // 字段注释
}

func genProtoFile(mp map[string]map[string]*fieldMeta) (data []string) {
	keys := maputil.Keys(mp)
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	for _, k := range keys {
		m := mp[k]
		metas := maputil.Values(m)
		slices.SortFunc(metas, func(a, b *fieldMeta) int {
			return cmp.Compare(a.Sort, b.Sort)
		})
		var fieldList []*protoField
		for _, meta := range metas {
			s := &protoField{
				fieldName: meta.ObjName,
				valueType: meta.DataType,
				mark:      meta.ObjDescribe,
			}
			// 转换PB字段类型
			if v, ok := xlsxType2ProtoTypeMap[s.valueType]; ok {
				s.valueType = v
			}
			// 枚举转换为int32
			if _, ok := GetXlsxStructHub().Enum[s.valueType]; ok {
				// 枚举修改为int32
				s.valueType = "int32"
			}

			if meta.DataSlicing != "" {
				// 是否为数组
				s.valueType = fmt.Sprintf("%s%s", "repeated ", s.valueType)
			}
			fieldList = append(fieldList, s)
		}
		data = append(data, createProtoMessage(k, fieldList)...)
	}
	return
}

func createProtoMessage(fileName string, fieldList []*protoField) []string {
	var data []string
	data = append(data, "//"+fileName)
	data = append(data, "message "+fileName+"\n{")
	for i, _data := range fieldList {
		str := fmt.Sprintf("\t%v %v = %v; //%v", _data.valueType, _data.fieldName, i+1, _data.mark)
		data = append(data, str)
	}
	data = append(data, "}")
	return data
}
