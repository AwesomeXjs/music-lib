package service

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type CustomClient struct {
	client *http.Client
}

type QueryParam struct {
	Key   string
	Value string
}

func (c *CustomClient) GetWithQuery(baseUrl, resource string, query ...QueryParam) (*http.Response, error) {
	inputFormat := func(str string) string {
		return strings.ReplaceAll(str, " ", "%20")
	}
	var queryString string

	for _, param := range query {
		queryString += fmt.Sprintf("%s=%s&", param.Key, inputFormat(param.Value))
	}
	url := baseUrl + resource + "?" + queryString
	req, err := c.client.Get(url)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Println("error happened", err)
		return nil, err
	}
	return req, nil
}

func NewClient() *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	return client
}
