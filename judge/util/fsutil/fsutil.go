package fsutil

import (
	"io"
	"os"
	"path"
	"runtime"
)

var judgeGoModDir string

func init() {
	_, testSourceFilePath, _, _ := runtime.Caller(0)
	judgeGoModDir = path.Clean(path.Join(
		path.Dir(testSourceFilePath),
		"..",
		"..",
	))
}

// go.mod のあるディレクトリの絶対パスを返す
func GetGoModAbsDir() string {
	return judgeGoModDir
}

func CopyFile(srcPath string, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
