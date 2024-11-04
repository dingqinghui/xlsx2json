# xlsx2json



## 说明

将配置的xlsx文件转为json格式，并自动生成Go代码解析json配置。支持自定义数据结构，自定义枚举。



## xlsx格式

参考examples



## 导表启动 

./bin/xlsx2json.exe -c config.json



## 启动配置格式

```go
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
```



## Go项目解析Json配置

```go
package examples

import (
	"github.com/dingqinghui/xlsx2json/bin/outDir/go/staticdata"
	"testing"
)

func TestXlsx2Json(t *testing.T) {
	StaticBeanData := new(staticdata.StaticBeanData)
	StaticBeanData.LoadAll("./output")
}

```



