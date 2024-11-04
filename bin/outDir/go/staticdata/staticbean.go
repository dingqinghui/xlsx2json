/**
 * @Author: dingQingHui
 * @Description:
 * @File: bean.tpl
 * @Version: 1.0.0
 * @Date: 2024/10/29 16:27
 */

package staticdata

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func jsonUnmarshal(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err2 := json.Unmarshal(data, v); err2 != nil {
		return err2
	}
	return nil
}

type StaticBeanData struct {
	TestTableDict  map[int32]*TestTable
	TestTableArray []*TestTable
}

func (v *StaticBeanData) LoadTestTable(path string) {
	fileName := filepath.Join(path, "TestTable.json")
	var TestTableArray []*TestTable
	if err := jsonUnmarshal(fileName, &TestTableArray); err != nil {
		return
	}
	var TestTableDict = make(map[int32]*TestTable)
	for _, bean := range TestTableArray {
		TestTableDict[bean.GetID()] = bean
	}
	v.TestTableDict = TestTableDict
	v.TestTableArray = TestTableArray
}

func (v *StaticBeanData) LoadAll(path string) {
	v.LoadTestTable(path)

}
