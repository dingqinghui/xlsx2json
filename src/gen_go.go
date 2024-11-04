/**
 * @Author: dingQingHui
 * @Description:
 * @File: gen_go
 * @Version: 1.0.0
 * @Date: 2024/10/29 16:43
 */

package src

import (
	"html/template"
	"log"
	"os"
)

type GenStruct struct {
	Name string
	Type string
}

type GenStructList struct {
	List []*GenStruct
}

func GenGoCode() {
	tmpl, err := template.ParseFiles("./bean.tpl")
	if err != nil {
		log.Fatal("template.ParseFiles", err)
		return
	}
	tplParam := new(GenStructList)
	for sheet, s := range GetXlsxStructHub().TableHeader {
		t := s["ID"].DataType
		if v, ok := xlsxType2ProtoTypeMap[t]; ok {
			t = v
		}
		sheet = CamelStr(sheet)
		tplParam.List = append(tplParam.List, &GenStruct{Name: sheet, Type: t})
	}
	_ = MkDirAll(tmpOutGoBeanFile)
	f, err := os.Create(tmpOutGoBeanFile)
	if err != nil {
		log.Fatal(tmpOutGoBeanFile, err)
		return
	}
	if err = tmpl.Execute(f, tplParam); err != nil {
		log.Fatal("tmpl.Execute", err)
		return
	}
	if err = f.Close(); err != nil {
		log.Fatal(tmpOutGoBeanFile, err)
		return
	}
	FormatGoFile(tmpOutGoBeanFile)

	log.Println("解析文件生成成功", tmpOutGoBeanFile)
}
