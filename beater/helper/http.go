package helper

import (
	"fmt"
	"github.com/wangdisdu/esbeat/config"
	"io/ioutil"
	"net/http"
)

type HTTP struct {
	client *http.Client
	config config.Config
}

// You can share it in diffrent instance
func NewHTTP(cfg config.Config) *HTTP {
	client := &http.Client{
		Timeout: cfg.Timeout,
	}
	return &HTTP{
		client: client,
		config: cfg,
	}
}

// FetchResponse fetches a response for the http request.
// It's important that resp.Body has to be closed if this method is used.
func (h *HTTP) FetchResponse(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var auth bool = h.config.Username != ""
	if auth {
		req.SetBasicAuth(h.config.Username, h.config.Password)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// FetchResponse fetches a array of byte for the http request.
// It will return error when http response code is any else 200.
func (h *HTTP) FetchContent(url string) ([]byte, error) {
	resp, err := h.FetchResponse(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP error code %d to request %s", resp.StatusCode, url)
	}

	return ioutil.ReadAll(resp.Body)
}
