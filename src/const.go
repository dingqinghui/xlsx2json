package src

import "path/filepath"

// GlobalTableName 全局定义表 表名
const (
	structMetaSheet = "@Type" // 表结构定义
	annotateTag     = "#"
)

// 列的种类
const (
	LineStruct = "结构"
	LineEnum   = "枚举"
	LineTable  = "表头"
)

var (
	tmpProtoFileName   = "staticstruct.proto"
	tmpOutDir          = "outDir/"
	tmpOutJsonDir      = filepath.Join(tmpOutDir, "json")
	tmpOutGoDir        = filepath.Join(tmpOutDir, "go")
	tmpOutProtoFile    = filepath.Join(tmpOutDir, "proto", tmpProtoFileName)
	tmpOutStaticData   = filepath.Join(tmpOutGoDir, "staticdata")
	tmpOutGoEnumFile   = filepath.Join(tmpOutStaticData, "staticenum.go")
	tmpOutGoBeanFile   = filepath.Join(tmpOutStaticData, "staticbean.go")
	tmpOutGoStructFile = filepath.Join(tmpOutStaticData, "staticstruct.go")
)
