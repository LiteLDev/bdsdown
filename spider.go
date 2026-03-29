package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func getSpiderHeaders() http.Header {
	header := http.Header{}
	header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	header.Set("Cache-Control", "max-age=0")
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
	header.Set("Sec-Fetch-Dest", "document")
	header.Set("Sec-Fetch-Mode", "navigate")
	header.Set("Sec-Fetch-Site", "none")
	header.Set("Sec-Fetch-User", "?1")
	header.Set("Upgrade-Insecure-Requests", "1")
	return header
}

var spiderClient *http.Client

func init() {
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	spiderClient = &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}
}

// FetchVersions fetches the versions of BDS and returns a map of platform and package link.
func FetchVersions(link string) (map[string]*url.URL, error) {

	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	request.Header = getSpiderHeaders()
	res, err := spiderClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch version with given url, code=%d", res.StatusCode)
	}
	var response struct {
		Result struct {
			Links []struct {
				DownloadType string `json:"downloadType"`
				DownloadURL  string `json:"downloadUrl"`
			} `json:"links"`
		} `json:"result"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	result := make(map[string]*url.URL)
	for _, item := range response.Result.Links {
		u, err := url.Parse(item.DownloadURL)
		if err != nil {
			return nil, fmt.Errorf("invalid download url for %s: %w", item.DownloadType, err)
		}
		result[item.DownloadType] = u
	}
	return result, nil
}
