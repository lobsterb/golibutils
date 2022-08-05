package utils

import (
	"archive/tar"
	"io"
	"os"
	"path"
)

func Decompress(tarFilePath string, dstDir string) error {
	srcFile, err := os.Open(tarFilePath)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	tr := tar.NewReader(srcFile)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		_, err = CreateDir(dstDir)
		if err != nil {
			return err
		}
		filename := path.Join(dstDir, hdr.Name)
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		_, err = io.Copy(file, tr)
		if err != nil {
			file.Close()
			return err
		}
		file.Close()
	}
	return nil
}
