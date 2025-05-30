package BA

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func DownloadFile(url string, outputPath string) error {
	client := &http.Client{
		Timeout: 2 * time.Second, // set your preferred timeout
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: status %s", resp.Status)
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func EnsureFile(filename string, url string) error {
	if FileExists(filename) {
		return nil
	}
	return DownloadFile(url, filename)
}
