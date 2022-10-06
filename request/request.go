package request

import (
	"net/http"
	"time"

	"github.com/Junkes887/translate/artifacts"
)

func Request(url string) *http.Response {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	artifacts.HandlerError(err)

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html")

	res, err := client.Do(req)
	artifacts.HandlerError(err)

	return res
}
