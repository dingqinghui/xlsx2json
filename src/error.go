/**
 * @Author: dingQingHui
 * @Description:
 * @File: error
 * @Version: 1.0.0
 * @Date: 2024/10/31 10:28
 */

package src

import "errors"

var (
	errSheetRowCnt  = errors.New("表格至少包含两行")
	errNotSheetMeta = errors.New("表格结构未定义")
)
