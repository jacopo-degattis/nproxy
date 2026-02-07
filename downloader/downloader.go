package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// This package has the role to download resource from online
// and save them to disk. It should be PROVIDER agnostic
type Downloader struct {
	Client              *http.Client
	DownloadDir         string
	MaximumDownloadPool int // for simultaneous downloads
}

// TODO: maybe add support for custom headers
// Also I delegate folder structure and folding to the caller instead of the downloader package
func (dw *Downloader) DownloadFrom(resUrl string, outputName string) error {
	rootDownloadLocation := dw.DownloadDir
	req, err := http.NewRequest("GET", resUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/140.0.0.0 Safari/537.36")

	if err != nil {
		return fmt.Errorf("downloader: %s", err.Error())
	}

	res, err := dw.Client.Do(req)
	if err != nil {
		return fmt.Errorf("downloader: %s", err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("downloader: request failed with status code %d", res.StatusCode)
	}

	out, err := os.OpenFile(rootDownloadLocation+outputName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("downloader: %s", err.Error())
	}

	copied, err := io.Copy(out, res.Body)

	resourceContent, err := strconv.Atoi(res.Header.Get("Content-Length"))
	if copied != int64(resourceContent) {
		return fmt.Errorf("downloader: unmatched 'content-type' and written size")
	}

	return nil
}
