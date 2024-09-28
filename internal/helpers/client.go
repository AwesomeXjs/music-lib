package helpers

import (
	"fmt"
	"github.com/AwesomeXjs/music-lib/configs"
	"github.com/AwesomeXjs/music-lib/pkg/logger"
	"net"
	"net/http"
	"time"
)

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

type CustomClient struct {
	Client         *http.Client
	Logger         logger.Logger
	SideServiceUrl string
}

type QueryParam struct {
	Key   string
	Value string
}

func (c *CustomClient) GetWithQuery(resource string, query ...QueryParam) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, c.SideServiceUrl+resource, nil)

	q := request.URL.Query()
	for i := range query {
		q.Add(query[i].Key, query[i].Value)
	}
	request.Header.Add("Content-Type", "application/json")
	request.URL.RawQuery = q.Encode()
	req, err := c.Client.Get(request.URL.String())

	if err != nil {
		c.Logger.Info(RESPONSE_PREFIX, err.Error())
		return nil, fmt.Errorf("%v", err)
	}
	return req, nil
}

func NewCustomClient(logger logger.Logger) *CustomClient {
	config := configs.New(logger)
	return &CustomClient{
		Client:         NewClient(),
		Logger:         logger,
		SideServiceUrl: config.SideServiceUrl,
	}
}
