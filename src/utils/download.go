package utils

import (
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
func DownloadVersion(version string, isPreview bool) error {
	bar := progressbar.NewOptions(
		-1,
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription("  "),
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
	_, err := DownloadFile(uri, bar)
	return err
}
