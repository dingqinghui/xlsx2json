/**
 * @Author: dingQingHui
 * @Description:
 * @File: config
 * @Version: 1.0.0
 * @Date: 2024/10/29 15:00
 */

package src

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	//
	// XlsxInputPath
	// @Description: xlsx文件路径，可配置多个路径/文件
	//
	XlsxInputPath []string
	//
	// JsonOutputPath
	// @Description: 转换后json文件输出路径
	//
	JsonOutputPath string
	//
	// GoOutputPath
	// @Description: 转换后go解析文件输出路径
	//
	GoOutputPath string
	//
	// ProtoOutputPath
	// @Description: 转换后proto文件输出路径
	//
	ProtoOutputPath string
}

var Cfg = new(Config)

func InitConfig(path string) {
	data, err := os.ReadFile(filepath.Join(path, "config.json"))
	if err != nil {
		log.Fatal("读取配置文件失败", path, err)
		return
	}
	if err1 := json.Unmarshal(data, Cfg); err1 != nil {
		log.Fatal("解析配置文件失败", path, err1)
		return
	}
}
