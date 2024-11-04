/**
 * @Author: dingQingHui
 * @Description:
 * @File: gen_json
 * @Version: 1.0.0
 * @Date: 2024/10/29 14:52
 */

package src

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"github.com/xuri/excelize/v2"
	"log"
	"path/filepath"
	"sync"
)

// GenJson
// @Description: 生成JSON文件
func GenJson() {
	var wait sync.WaitGroup
	doFunc := func(path string) {
		WrapWaitGroup(&wait, func() {
			xlsxSheet2Json(path, &wait, false)
		})
	}
	TraversalPath(doFunc, 2, Cfg.XlsxInputPath...)
	wait.Wait()

}

func xlsxSheet2Json(filePath string, wait *sync.WaitGroup, escapeChar bool) {
	if !isXlsxFile(filePath) {
		return
	}
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("xlsxSheet2Json", filePath, err)
		return
	}
	// 遍历所有表单
	for _, sheetStr := range xlsx.GetSheetList() {
		sheet := sheetStr
		// 跳过注释表和定义表
		if isAnnotateTag(sheetStr) || isStructMetaSheet(sheet) {
			continue
		}

		WrapWaitGroup(wait, func() {
			sheetToJsonFile(xlsx, sheet, tmpOutJsonDir, escapeChar)
		})
	}
	return
}

func sheetToJsonFile(xlsx *excelize.File, sheet string, outputPath string, escapeChar bool) {
	rows, _ := xlsx.GetRows(sheet)
	if len(rows) <= 2 {
		log.Fatal(sheet, errSheetRowCnt)
		return
	}
	// 表格第一行，第一列元素，以#开头，跳过该表格
	if isAnnotateTag(rows[0][0]) {
		return
	}
	jsonObj := gabs.New()
	array, _ := jsonObj.Array()
	fieldDict := getStructFieldDict(sheet, rows[0])
	for i, rowData := range rows[2:] {
		row, have := tableRow2Json(sheet, i, rowData, fieldDict)
		if !have {
			continue
		}
		if err := array.ArrayAppend(row); err != nil {
			log.Fatalf("%v row:%v err:%v\n", sheet, i, err)
			return
		}
	}
	data := jsonObj.EncodeJSON(gabs.EncodeOptIndent("", "\t"))
	filename := filepath.Join(outputPath, fmt.Sprintf("%v.json", sheet))
	if err := WriteFile(filename, data); err != nil {
		log.Fatal(filename, err)
		return
	}
	log.Println("json文件生成成功", outputPath)
}

func tableRow2Json(sheet string, row int, rowData []string, fieldDict map[int]*fieldMeta) (*gabs.Container, bool) {
	have := false
	gObj := gabs.New()
	for idx, colCell := range rowData {
		// 跳过注释
		if colCell != "" && (isAnnotateTag(colCell) && idx == 0) {
			break
		}
		columnMeta, ok := fieldDict[idx]
		if !ok {
			continue
		}
		v := Cast(columnMeta, colCell)
		if _, err := gObj.SetP(v, columnMeta.ObjName); err != nil {
			log.Fatalf("列转换失败 %v row:%v column：%v data:%v err:%v\n", sheet, row, idx, colCell, err)
			return nil, false
		}
		have = true
	}
	return gObj, have
}

func getStructFieldDict(sheet string, firstRow []string) map[int]*fieldMeta {
	structData, ok := GetXlsxStructHub().TableHeader[sheet]
	if !ok {
		log.Fatal(sheet, errNotSheetMeta)
	}
	ret := make(map[int]*fieldMeta)
	ret[0] = structData["ID"]
	// 拼接第一行字符串
	//idx := 1
	for idx, v := range firstRow[1:] {
		if v == "" || isAnnotateTag(v) {
			// 跳过备注
			continue
		}
		if _, ok := structData[v]; !ok {
			// 跳过未定义类型的字段
			continue
		}
		ret[idx+1] = structData[v]
	}
	return ret
}
