/**
 * @Author: dingQingHui
 * @Description:
 * @File: func
 * @Version: 1.0.0
 * @Date: 2024/10/30 11:30
 */

package src

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"log"
	"path/filepath"
	"sync"
)

func ClearDir() {
	if err := EmptyDir(tmpOutJsonDir); err != nil {
		log.Fatal("EmptyDir", err)
		return
	}
	log.Println("清空文件夹", tmpOutJsonDir)
}

func Gen() {
	wait := &sync.WaitGroup{}
	WrapWaitGroup(wait, func() {
		GenJson()
	})

	WrapWaitGroup(wait, func() {
		GenEnumGoCode()
	})

	WrapWaitGroup(wait, func() {
		GenProtoFileCode()
	})

	wait.Wait()

	GenGoCode()

	GenProto2Go(tmpOutProtoFile, tmpOutGoStructFile)

}

func CopyCodeDir() {

	if Cfg.GoOutputPath != "" {
		_ = MkDirAll(Cfg.GoOutputPath)
		if err := fileutil.CopyDir(tmpOutGoDir, Cfg.GoOutputPath); err != nil {
			log.Fatal("CopyCodeDir", err)
			return
		}
		log.Println("拷贝文件成功", tmpOutGoDir, Cfg.GoOutputPath)
	}

	if Cfg.JsonOutputPath != "" {
		_ = MkDirAll(Cfg.JsonOutputPath)
		if err := fileutil.CopyDir(tmpOutJsonDir, Cfg.JsonOutputPath); err != nil {
			log.Fatal("CopyCodeDir", err)
			return
		}
		log.Println("拷贝文件成功", tmpOutJsonDir, Cfg.JsonOutputPath)
	}

	path := filepath.Join(Cfg.ProtoOutputPath, tmpProtoFileName)
	_ = MkDirAll(path)
	if err := fileutil.CopyFile(tmpOutProtoFile, path); err != nil {
		log.Fatal("CopyCodeDir", err)
		return
	}
	log.Println("拷贝文件成功", tmpOutProtoFile, path)
}
