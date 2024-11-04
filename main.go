package main

import (
	"flag"
	"log"
	. "xlsx2json/src"
)

func main() {
	path := flag.String("c", "./config.json", "config path")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("config path:%v\n", *path)
	// 解析配置文件
	InitConfig(*path)
	// 清空临时文件夹
	ClearDir()
	// 读取所有表头结构
	GetXlsxStructHub().Do()
	// 生成文件
	Gen()
	// 复制生成文件到指定目录
	CopyCodeDir()
}
