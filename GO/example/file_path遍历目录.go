package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	_ = os.RemoveAll("./cache/web.zip")
	file, _ := os.Create("./cache/web.zip")
	defer file.Close()
	zipFile := zip.NewWriter(file)
	defer zipFile.Close()

	_ = filepath.Walk("/Users/harder/github.com-codes/bups/example", func(path string, info os.FileInfo, _ error) error {
		fmt.Println(path)
		if path == "/Users/harder/github.com-codes/bups/example" {
			return nil
		}
		// 获取头文件信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, "/Users/harder/github.com-codes/bups/example"+"/")
		// 设置压缩算法
		header.Method = zip.Deflate
		// 创建压缩包头部信息
		w, _ := zipFile.CreateHeader(header)
		newfile, _ := os.Open(path)
		defer newfile.Close()
		_, _ = io.Copy(w, newfile)
		return nil
	})
}
