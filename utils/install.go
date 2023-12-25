package utils

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/schollz/progressbar/v3"
)

// Unzip zip file to the current directory.
func Unzip(zipFile *os.File, bar *progressbar.ProgressBar) error {
	fileStat, err := zipFile.Stat()
	if err != nil {
		return err
	}

	excludedFilesMap := make(map[string]bool)
	for _, file := range GetConfig().ExcludedFiles {
		excludedFilesMap[file] = true
	}

	reader, err := zip.NewReader(zipFile, fileStat.Size())
	if err != nil {
		return err
	}
	reader.RegisterDecompressor(zip.Deflate, flate.NewReader)

	bar.ChangeMax64(fileStat.Size())
	for _, file := range reader.File {
		if excludedFilesMap[file.Name] {
			bar.Add64(int64(file.CompressedSize64))
			continue
		}

		if file.FileInfo().IsDir() {
			os.Mkdir(file.Name, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		fileWriter, err := os.OpenFile(file.Name, os.O_CREATE|os.O_WRONLY, file.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			return err
		}

		bar.Add64(int64(file.CompressedSize64))
	}
	bar.Finish()

	return nil
}

func CheckCache(version string, cacheDir string) (string, error) {
	target := path.Join(cacheDir, "bedrock-server-"+version+".zip")
	_, err := os.Stat(target)
	if err == nil {
		return target, nil
	}
	return "", err
}

// Install installs the given version of BDS.
func Install() error {
	version := GetConfig().TargetVersion
	usePreview := GetConfig().UsePreview
	useCache := GetConfig().UseCache
	cacheDir := GetConfig().CacheDir

	var path string
	var err error

	if useCache {
		fmt.Println("Checking cache...")
		path, err = CheckCache(version, cacheDir)
		if err == nil {
			fmt.Println(" Found cache!")
			fmt.Println("Unziping cached files...")
			goto Unzip
		} else {
			fmt.Println(" Cache not found.")
		}
	}

	fmt.Println("Downloading BDS v" + version + "...")
	path, err = DownloadVersion(version, usePreview)
	if err != nil {
		return err
	}
	fmt.Println(" Download complete!")

	fmt.Println("Unziping downloaded files...")

Unzip:
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return err
	}

	bar := progressbar.NewOptions(
		114514,
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription(" Unziping...   "),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	Unzip(file, bar)
	fmt.Println(" Unzip complete!")

	file.Close()
	if !useCache {
		fmt.Println("Cleaning up...")
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}
