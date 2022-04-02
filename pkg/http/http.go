package http

import (
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var defaultTimeout = 30 * time.Second

var client *resty.Client

func NewHttpClient(timeout time.Duration) *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		DisableKeepAlives:     false,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          500,
		MaxIdleConnsPerHost:   500,
		MaxConnsPerHost:       2000,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

func InitClient(t string, debug bool) {
	timeout, err := time.ParseDuration(t)
	if err != nil || timeout.Seconds() == 0 {
		timeout = defaultTimeout
	}

	client = resty.NewWithClient(NewHttpClient(timeout)).SetTimeout(timeout).SetDebug(debug)
}
