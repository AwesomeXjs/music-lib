package service

import (
	"fmt"
	"net"
	"net/http"
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
	request, err := http.NewRequest(http.MethodGet, baseUrl+resource, nil)

	q := request.URL.Query()
	for i := range query {
		q.Add(query[i].Key, query[i].Value)
		fmt.Println(query[i].Key, query[i].Value)
	}
	request.Header.Add("Content-Type", "application/json")

	request.URL.RawQuery = q.Encode()
	req, err := c.client.Get(request.URL.String())

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
