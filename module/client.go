package module

import (
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36"
	ConnNum   = 10
)

var client = NewHTTPClient(ConnNum)

func NewHTTPClient(connNum int) *http.Client {
	jar, _ := cookiejar.New(nil)

	tp := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxConnsPerHost:       connNum,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	c := &http.Client{
		Jar:       jar,
		Transport: tp,
	}
	return c
}

func FileSize(url string) (int, error) {
	req, err := NewHTTPReqHead(url)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	size := resp.Header.Get("Content-Length")
	return strconv.Atoi(size)
}

func NewHTTPReqGet(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}

func NewHTTPReqHead(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", UserAgent)
	return req, nil
}

func SetHTTPReqHeaderRange(url string, start, end int64) (*http.Request, error) {
	req, err := NewHTTPReqGet(url)
	if err != nil {
		return nil, err
	}

	Range := fmt.Sprintf("bytes=%d-%d", start, end)
	req.Header.Set("Range", Range)
	return req, nil
}
