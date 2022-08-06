package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Unzip zip解压
func Unzip(zipFile string, dstDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		// 解压后路径
		outFilePath := filepath.Join(dstDir, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(outFilePath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(outFilePath), os.ModePerm); err != nil {
				return err
			}
			inFile, err := file.Open()
			if err != nil {
				return err
			}
			outFile, err := os.Create(outFilePath)
			if err != nil {
				inFile.Close()
				return err
			}
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				inFile.Close()
				outFile.Close()
				return err
			}
			inFile.Close()
			outFile.Close()
		}
	}
	return nil
}
