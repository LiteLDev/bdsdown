package main

import (
	"archive/zip"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
	"path/filepath"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func UnzipPackage(file, dest string) error {
	log.Infof("Extracting %s to %s", file, dest)
	zipFile, err := zip.OpenReader(file)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	bar := pb.StartNew(len(zipFile.File))
	os.MkdirAll(dest, 644)

	for _, f := range zipFile.File {
		extractFile := func(path string, file *zip.File) error {
			if fileExists(path) && config.InstallerExcludeSet.Contains(f.Name) {
				return nil
			}

			fileWriter, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, f.Mode())
			if err != nil {
				return err
			}
			defer fileWriter.Close()

			fileReader, err := f.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()

			_, err = io.Copy(fileWriter, fileReader)
			if err != nil {
				return err
			}
			return nil
		}
		bar.Add(1)
		realPath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(realPath, f.Mode())
		} else {
			extractFile(realPath, f)
		}

	}
	bar.Finish()
	return nil
}
