package bitaxe

import (
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	SystemInfo(ctx context.Context, baseAddress string) (SystemInfo, error)
}

func NewClient(httpClient *http.Client) Client {
	httpClient = cmp.Or(httpClient, &http.Client{
		Timeout: 5 * time.Second,
	})

	return &client{
		httpClient: httpClient,
	}
}

type client struct {
	httpClient *http.Client
}

// The ESP-miner code has more details for each endpoint
//
// https://github.com/skot/ESP-Miner/blob/master/main/http_server/http_server.c

func (c *client) SystemInfo(ctx context.Context, baseAddress string) (SystemInfo, error) {
	var info SystemInfo

	resp, err := c.httpClient.Get(baseAddress + "/api/system/info")
	if err != nil {
		return info, fmt.Errorf("get system info: %w", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return info, fmt.Errorf("reading system info json: %w", err)
	}

	return info, nil
}
