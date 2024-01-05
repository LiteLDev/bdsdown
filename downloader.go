package main

import (
	"context"
	"github.com/LiteLDev/pget"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func newDownloadClient(maxIdleConnsPerHost int) *http.Client {
	tr := http.DefaultTransport.(*http.Transport).Clone()
	tr.MaxIdleConns = 0 // no limit
	tr.MaxIdleConnsPerHost = maxIdleConnsPerHost
	tr.DisableCompression = true
	return &http.Client{
		Transport: tr,
	}
}

func DownloadFile(url, dest string) error {
	client := newDownloadClient(16)

	target, err := pget.Check(context.Background(), &pget.CheckConfig{
		URLs:    []string{url},
		Timeout: 10 * time.Second,
		Client:  client,
	})
	if err != nil {
		return err
	}

	dir, _ := filepath.Split(dest)
	os.MkdirAll(dir, 644)

	opts := []pget.DownloadOption{
		pget.WithUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36", ""),
		pget.WithReferer(""),
	}

	err = pget.Download(context.Background(), &pget.DownloadConfig{
		Filename:      target.Filename,
		Dirname:       dir,
		ContentLength: target.ContentLength,
		Procs:         config.Procs,
		URLs:          target.URLs,
		Client:        client,
	}, opts...)
	return err
}
