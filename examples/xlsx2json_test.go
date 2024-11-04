/**
 * @Author: dingQingHui
 * @Description:
 * @File: xlsx2json_test
 * @Version: 1.0.0
 * @Date: 2024/11/4 10:36
 */

package examples

import (
	"github.com/dingqinghui/xlsx2json/bin/outDir/go/staticdata"
	"testing"
)

func TestXlsx2Json(t *testing.T) {
	StaticBeanData := new(staticdata.StaticBeanData)
	StaticBeanData.LoadAll("./output")
}
