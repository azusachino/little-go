package error

import (
	"fmt"
	"io"
	"os"
)

// 打开资源的同时考虑到关闭资源
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// recover的使用场景
func ParseJSON(input string) (s string, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("JSON:internal error: %v", p)
		}
	}()
	return input, err
}
