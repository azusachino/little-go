package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5(data []byte) []byte {
	digest := md5.New()
	digest.Write(data)
	return digest.Sum(nil)
}

func Md5Hex(data []byte) string {
	return fmt.Sprintf("%x", Md5(data))
}
