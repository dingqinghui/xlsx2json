/**
 * @Author: dingQingHui
 * @Description:
 * @File: gen_enum
 * @Version: 1.0.0
 * @Date: 2024/10/28 18:33
 */

package src

import (
	"cmp"
	"fmt"
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/spf13/cast"
	"log"
	"slices"
	"strings"
)

var (
	enumHead = `
	// auto gen don't edit
	package staticdata

	`
)

func GenEnumGoCode() {
	var data []string
	data = append(data, enumHead)

	enums := GetXlsxStructHub().Enum
	keys := maputil.Keys(enums)
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	for _, k := range keys {
		m := enums[k]
		metas := maputil.Values(m)
		slices.SortFunc(metas, func(a, b *fieldMeta) int {
			return cmp.Compare(cast.ToInt32(a.DataDefault), cast.ToInt32(b.DataDefault))
		})

		data = append(data, "const (")
		for _, meta := range metas {
			if meta.DataDefault == "" {
				continue
			}
			data = append(data, fmt.Sprintf("%v_%v %v = %v // %v", meta.ObjType, meta.ObjName, meta.DataType, meta.DataDefault, meta.ObjDescribe))
		}
		data = append(data, ")")
	}

	_ = WriteFile(tmpOutGoEnumFile, []byte(strings.Join(data, "\n")))

	FormatGoFile(tmpOutGoEnumFile)

	log.Println("枚举文件生成成功", tmpOutGoEnumFile)
}
