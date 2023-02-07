package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/schollz/progressbar/v3"
)

// DownloadFile downloads a file from the given uri and returns the file name.
// If bar is not nil, the progress will be written to it.
func DownloadFile(uri string, bar *progressbar.ProgressBar) (string, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download file %s (%s)", uri, res.Status)
	}

	fileName := path.Base(res.Request.URL.Path)

	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	if bar != nil {
		bar.ChangeMax(int(res.ContentLength))
		_, err = io.Copy(io.MultiWriter(file, bar), res.Body)
	} else {
		_, err = io.Copy(file, res.Body)
	}

	if err != nil {
		return "", err
	}

	return fileName, nil
}

// DownloadVersion downloads the given version of BDS.
func DownloadVersion(version string, isPreview bool) (string, error) {
	bar := progressbar.NewOptions(
		114514,
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription(" Downloading..."),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	var uri string
	if isPreview {
		uri = PreviewDownloadPrefix + version + PreviewDownloadSuffix
	} else {
		uri = ReleaseDownloadPrefix + version + ReleaseDownloadSuffix
	}
	return DownloadFile(uri, bar)
}
