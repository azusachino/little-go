package fs

import (
	"github.com/azusachino/golong/tools/hash"
	"os"
)

// TempFileWithText The caller has responsibility to close the fd and delete file with name
func TempFileWithText(text string) (*os.File, error) {
	tmpFile, err := os.CreateTemp(os.TempDir(), hash.Md5Hex([]byte(text)))
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(tmpFile.Name(), []byte(text), os.ModeTemporary); err != nil {
		return nil, err
	}
	return tmpFile, nil
}
