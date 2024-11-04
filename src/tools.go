/**
 * @Author: dingQingHui
 * @Description:
 * @File: tools
 * @Version: 1.0.0
 * @Date: 2024/10/29 14:55
 */

package src

import (
	"fmt"
	"github.com/duke-git/lancet/v2/system"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Itoa(num interface{}) string {
	switch n := num.(type) {
	case int8:
		return strconv.FormatInt(int64(n), 10)
	case int16:
		return strconv.FormatInt(int64(n), 10)
	case int32:
		return strconv.FormatInt(int64(n), 10)
	case int:
		return strconv.FormatInt(int64(n), 10)
	case int64:
		return strconv.FormatInt(n, 10)
	case uint8:
		return strconv.FormatUint(uint64(n), 10)
	case uint16:
		return strconv.FormatUint(uint64(n), 10)
	case uint32:
		return strconv.FormatUint(uint64(n), 10)
	case uint:
		return strconv.FormatUint(uint64(n), 10)
	case uint64:
		return strconv.FormatUint(n, 10)
	case string:
		return num.(string)
	}
	return ""
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// EmptyDir 清空指定文件夹
func EmptyDir(path string) error {
	dirPath := filepath.Join(path)

	d, err := os.Open(dirPath)
	if err != nil {
		return nil
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dirPath, name))
		if err != nil {
			return err
		}
	}

	return nil
}

func WrapWaitGroup(wait *sync.WaitGroup, f func()) {
	wait.Add(1)
	go func() {
		f()
		wait.Done()
	}()
}

func TraversalPath(f func(path string), depth int, paths ...string) {
	for _, path := range paths {
		if isFile(path) {
			f(path)
		} else {
			readDir(path, 1, depth, f)
		}
	}
}

func readDir(filePath string, depth, maxDepth int, f func(path string)) {
	if depth >= maxDepth {
		return
	}
	rd, err := os.ReadDir(filePath)
	if err != nil {
		log.Fatal("不存在的目录或文件", err)
		return
	}
	// 读取表格数据
	for _, fi := range rd {
		fileName := fi.Name()
		if fi.IsDir() {
			readDir(filepath.Join(filePath, fileName), depth+1, maxDepth, f)
		} else {
			f(filepath.Join(filePath, fileName))
		}
	}
}

// TimeCost 函数耗时检测
func TimeCost(src string) func() {
	start := time.Now()
	return func() {
		log.Println(src, time.Since(start))
	}
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		log.Fatal("isFile", path, err)
		return false
	}
	return info.Mode().IsRegular()
}

func FormatGoFile(path string) {
	var args []string
	args = append(args, "fmt")
	args = append(args, path)

	stdout, stderr, err := system.ExecCommand("go", func(cmd *exec.Cmd) {
		cmd.Args = append(cmd.Args, args...)
	})
	if stderr != "" || err != nil {
		log.Fatal("FormatGoFile", stdout, stderr, err)
	}

}

func GenProto2Go(protoPath, goPath string) {
	_ = MkDirAll(goPath)

	dir := filepath.Dir(protoPath)
	var args []string
	args = append(args, fmt.Sprintf("--proto_path=%v", dir))
	args = append(args, fmt.Sprintf("--go_out=%v", filepath.Dir(goPath)))
	args = append(args, protoPath)
	stdout, stderr, err := system.ExecCommand("./tools/bin/protoc.exe", func(cmd *exec.Cmd) {
		cmd.Args = append(cmd.Args, args...)
	})

	if stderr != "" || err != nil {
		log.Fatal("GenProto2Go", stdout, stderr, err)
	}

	log.Println("proto转Go生成成功", protoPath, goPath)
}

func MkDirAll(filePath string) error {
	// 获取目录路径
	dirPath := filepath.Dir(filePath)
	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，递归创建
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatalf("MkdirAll path:%v err:%v\n", dirPath, err)
			return err
		}
	}
	return nil
}

func WriteFile(filePath string, data []byte) error {
	// 获取目录路径
	dirPath := filepath.Dir(filePath)
	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，递归创建
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			log.Fatalf("MkdirAll path:%v err:%v\n", dirPath, err)
			return err
		}
	}
	filepath.Dir(filePath)
	// 目录存在或创建成功，可以进行文件写入
	if err := os.WriteFile(filePath, data, 0666); err != nil {
		log.Fatalf("WriteFile  path:%v err:%v\n", filePath, err)
		return err
	}
	return nil
}

func isXlsxFile(path string) bool {
	return strings.HasSuffix(path, "xlsx")
}

func isAnnotateTag(s string) bool {
	return strings.HasPrefix(s, annotateTag)
}

func isStructMetaSheet(s string) bool {
	return strings.HasPrefix(s, structMetaSheet)
}
