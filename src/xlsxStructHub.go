// Package src
// @Description: 加载所有的表头的结构体
package src

import (
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// 全局定义表首行类型
const (
	ValueType   = "种类"
	ObjType     = "对象类型"
	ObjDescribe = "中文描述"
	ObjName     = "字段名"
	DataType    = "字段类型"
	DataSlicing = "数组切割"
	DataDefault = "默认值"
	Filtrate    = "筛选"
)

// 全局定义表结构
type fieldMeta struct {
	ValueType   string `xlsx:"种类"`
	ObjType     string `xlsx:"对象类型"`
	ObjDescribe string `xlsx:"中文描述"`
	ObjName     string `xlsx:"字段名"`
	DataType    string `xlsx:"字段类型"`
	DataSlicing string `xlsx:"数组切割"`
	DataDefault string `xlsx:"默认值"`
	Filtrate    string `xlsx:"筛选"`
	Sort        int
}

func (g *fieldMeta) IsSlice() bool {
	return g.DataSlicing != ""
}
func (g *fieldMeta) Copy() *fieldMeta {
	c := *g
	return &c
}

var (
	xlsxHub     *XlsxStructHub
	xlsxHubOnce sync.Once
)

func GetXlsxStructHub() *XlsxStructHub {
	xlsxHubOnce.Do(func() {
		xlsxHub = &XlsxStructHub{
			Enum:        make(map[string][]*fieldMeta),
			Struct:      make(map[string]map[string]*fieldMeta),
			TableHeader: make(map[string]map[string]*fieldMeta),
			dataChan:    make(chan *fieldMeta),
			wg:          &sync.WaitGroup{},
		}
	})
	return xlsxHub
}

type XlsxStructHub struct {
	Enum        map[string][]*fieldMeta          // 所有的枚举
	Struct      map[string]map[string]*fieldMeta // 所有的结构体
	TableHeader map[string]map[string]*fieldMeta // 所有的表头都定义
	dataChan    chan *fieldMeta
	wg          *sync.WaitGroup
	tableTitle  map[string]string // map[表名]map[字段名]{ 类型，中文描述 }
}

func (x *XlsxStructHub) Do() {
	paths := Cfg.XlsxInputPath
	// 读取各个路径的配置表
	for _, path := range paths {
		if isFile(path) {
			WrapWaitGroup(x.wg, func() {
				x.readTableStruct(path)
			})
			WrapWaitGroup(x.wg, func() {
				x.readTableIdType(path)
			})
		} else {
			x.readDir(path)
		}
	}

	go x.receiveChan()

	// 等待所有协程读取完数据
	x.wg.Wait()
	// 关闭管道，读取数据结束
	close(x.dataChan)

	log.Printf("完成xlsx结构读取")
}

// 读取全局配置的表单数据
func (x *XlsxStructHub) readDir(filePath string) {
	rd, err := os.ReadDir(filePath)
	if err != nil {
		log.Println("不存在的目录或文件", err)
		return
	}
	// 读取表格数据
	for _, fi := range rd {
		FileName := fi.Name()
		if fi.IsDir() {
			// 跳过文件夹
			continue
		} else {
			WrapWaitGroup(x.wg, func() {
				x.readTableStruct(filepath.Join(filePath, FileName))
			})
			WrapWaitGroup(x.wg, func() {
				x.readTableIdType(filepath.Join(filePath, FileName))
			})
		}
	}
}

func (x *XlsxStructHub) receiveChan() {
	for cv := range x.dataChan {
		switch cv.ValueType {
		case LineEnum:
			// 枚举
			// XlsxStructHub.keyType["EItemType"] = row["枚举"]  or XlsxStructHub.keyType["ItemConfig"] = row["结构"]
			//if _, ok := x.Enum[cv.ObjType]; !ok {
			//	x.Enum[cv.ObjType] = make(map[string]*fieldMeta)
			//}
			x.Enum[cv.ObjType] = append(x.Enum[cv.ObjType], cv)
		case LineStruct:
			// 结构,表头
			if _, ok := x.Struct[cv.ObjType]; !ok {
				x.Struct[cv.ObjType] = make(map[string]*fieldMeta)
			}
			x.Struct[cv.ObjType][cv.ObjDescribe] = cv
		case LineTable:
			// 表头
			if _, ok := x.TableHeader[cv.ObjType]; !ok {
				x.TableHeader[cv.ObjType] = make(map[string]*fieldMeta)
			}
			x.TableHeader[cv.ObjType][cv.ObjName] = cv
		}
	}
}

// 读取全局配置的表单数据
func (x *XlsxStructHub) readTableStruct(filepath string) {
	if !isXlsxFile(filepath) {
		return
	}
	// 打开xlsx文件
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Fatal("打开文件失败", filepath, err)
		return
	}
	var firstRow []string
	// 获取所有的表单
	for _, sheetName := range f.GetSheetList() {

		// @Type开头
		if !isStructMetaSheet(sheetName) {
			continue
		}
		// 读取表单数据
		rows, err1 := f.GetRows(sheetName)
		if err1 != nil {
			log.Fatal("打开文件失败", filepath, sheetName, err1)
			return
		}
		firstRow = make([]string, 0)
		// 读取第一行，数据类型
		for _, value := range rows[0] {
			firstRow = append(firstRow, value)
		}
		var sort int = 1
		for _, row := range rows[1:] {
			// 读取一整行数据
			field := genFieldMeta(firstRow, row)
			field.Sort = sort
			x.dataChan <- field
			sort++
		}
	}
}

func (x *XlsxStructHub) readTableIdType(filepath string) {
	if !isXlsxFile(filepath) {
		return
	}
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		log.Fatal("OpenFile", filepath, err)
		return
	}
	for _, sheetStr := range xlsx.GetSheetList() {
		sheet := sheetStr
		// 跳过#开头的表， 跳过定义
		if isAnnotateTag(sheetStr) {
			continue
		}

		if isStructMetaSheet(sheetStr) {
			continue
		}
		x.readSheetHead(xlsx, sheet)
	}
}
func (x *XlsxStructHub) readSheetHead(xlsx *excelize.File, sheet string) {
	rows, _ := xlsx.GetRows(sheet)
	if len(rows) <= 2 {
		return
	}
	// 表格第一行，第一列元素，以#开头，跳过该表格
	if isAnnotateTag(rows[0][0]) {
		return
	}
	t := "int"
	if !isAnnotateTag(rows[1][0]) && rows[1][0] != "" {
		switch rows[1][0] {
		case "string":
			t = rows[1][0]
		}
	}

	idField := &fieldMeta{
		ValueType:   LineTable,
		ObjType:     sheet,
		ObjDescribe: "ID",
		ObjName:     "ID",
		DataType:    t,
		Sort:        0,
	}
	x.dataChan <- idField
}

func (x *XlsxStructHub) GetMeta(structName, fieldName string) *fieldMeta {
	sMeta := x.Struct[structName]
	if sMeta == nil {
		return nil
	}
	for _, meta := range sMeta {
		if meta.ObjDescribe == fieldName {
			return meta
		}
		if meta.ObjName == fieldName {
			return meta
		}
	}
	return nil
}
func (x *XlsxStructHub) GetEnumMeta(enumName string, enumValue string) (*fieldMeta, bool) {
	if enumMap, ok := x.Enum[enumName]; ok {
		for _, meta := range enumMap {
			if meta.ObjDescribe == enumValue {
				return meta, ok
			}
			if meta.ObjName == enumValue {
				return meta, ok
			}
		}
		return nil, ok
	} else {
		return nil, ok
	}

}

func genFieldMeta(convert, row []string) *fieldMeta {
	ret := &fieldMeta{}
	for idx, cell := range row {
		if idx >= len(convert) {
			return ret
		}
		switch convert[idx] {
		case ValueType:
			ret.ValueType = cell
			break
		case ObjType:
			ret.ObjType = cell
			break
		case ObjDescribe:
			ret.ObjDescribe = cell
			break
		case ObjName:
			ret.ObjName = cell
			break
		case DataType:
			ret.DataType = cell
			break
		case DataSlicing:
			ret.DataSlicing = cell
			break
		case DataDefault:
			ret.DataDefault = cell
			break
		case Filtrate:
			ret.Filtrate = cell
			break
		default:
			//log.Println("全局定义表，未支持的列，如有需求请联系程序", convert[idx], cell)
		}
	}
	return ret
}
