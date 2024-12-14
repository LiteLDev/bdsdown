package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(url, dest string) error {
	// If there's a file, don't download it twice
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	dir, _ := filepath.Split(dest)
	os.MkdirAll(dir, 0755)
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Get file size
	size := resp.ContentLength
	progress := &ProgressReader{Reader: resp.Body, Total: size}

	// Enable progressbar goroutine
	go progress.PrintProgress()

	_, err = io.Copy(out, progress)
	if err != nil {
		return err
	}

	return nil
}

type ProgressReader struct {
	io.Reader
	Total   int64
	Current int64
}

func (pr *ProgressReader) Read(p []byte) (int, error) {
	n, err := pr.Reader.Read(p)
	pr.Current += int64(n)
	return n, err
}

func (pr *ProgressReader) PrintProgress() {
	for {
		percentage := float64(pr.Current) / float64(pr.Total) * 100
		log.Infof("Downloading... %.2f%% complete", percentage)
		if pr.Current >= pr.Total {
			log.Info("Download complete")
			break
		}
		time.Sleep(1 * time.Second)
	}
}
