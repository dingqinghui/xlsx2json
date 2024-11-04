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
	XlsxInputPath   []string
	JsonOutputPath  string
	GoOutputPath    string
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
