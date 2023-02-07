package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	DownloadPage          = "https://www.minecraft.net/en-us/download/server/bedrock"
	ReleaseDownloadPrefix = "https://minecraft.azureedge.net/bin-win/bedrock-server-"
	ReleaseDownloadSuffix = ".zip"
	PreviewDownloadPrefix = "https://minecraft.azureedge.net/bin-win-preview/bedrock-server-"
	PreviewDownloadSuffix = ".zip"
)

func fetchDownloadPage() (string, error) {
	header := make(http.Header)
	header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	header.Set("Cache-Control", "max-age=0")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	header.Set("Sec-Fetch-Dest", "document")
	header.Set("Sec-Fetch-Mode", "navigate")
	header.Set("Sec-Fetch-Site", "none")
	header.Set("Sec-Fetch-User", "?1")
	header.Set("Upgrade-Insecure-Requests", "1")

	req, err := http.NewRequest("GET", DownloadPage, nil)
	if err != nil {
		return "", err
	}
	req.Header = header
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch download page: %s", res.Status)
	}

	defer res.Body.Close()

	buf := make([]byte, 1024)
	var sb strings.Builder
	for {
		n, err := res.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		sb.Write(buf[:n])
	}

	return sb.String(), nil
}

// GetLatestReleaseVersion returns the latest version of BDS.
func GetLatestReleaseVersion() (string, error) {
	body, err := fetchDownloadPage()
	if err != nil {
		return "", err
	}

	prefixIndex := strings.Index(body, ReleaseDownloadPrefix)
	if prefixIndex == -1 {
		return "", fmt.Errorf("failed to find release download prefix")
	}
	body = body[prefixIndex+len(ReleaseDownloadPrefix):]
	suffixIndex := strings.Index(body, ReleaseDownloadSuffix)
	if suffixIndex == -1 {
		return "", fmt.Errorf("failed to find release download suffix")
	}
	version := body[:suffixIndex]

	return version, nil
}

// GetLatestPreviewVersion returns the latest preview version of BDS.
func GetLatestPreviewVersion() (string, error) {
	body, err := fetchDownloadPage()
	if err != nil {
		return "", err
	}

	prefixIndex := strings.Index(body, PreviewDownloadPrefix)
	if prefixIndex == -1 {
		return "", fmt.Errorf("failed to find preview download prefix")
	}
	body = body[prefixIndex+len(PreviewDownloadPrefix):]
	suffixIndex := strings.Index(body, PreviewDownloadSuffix)
	if suffixIndex == -1 {
		return "", fmt.Errorf("failed to find preview download suffix")
	}
	version := body[:suffixIndex]

	return version, nil
}
