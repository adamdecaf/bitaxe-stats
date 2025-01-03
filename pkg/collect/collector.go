package collect

import (
	"context"
	"fmt"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
)

type Collector interface {
	SystemInfo(ctx context.Context, targets []string) ([]bitaxe.SystemInfo, error)
}

func NewCollector(client bitaxe.Client) Collector {
	return &collector{
		client: client,
	}
}

type collector struct {
	client bitaxe.Client
}

func (c *collector) SystemInfo(ctx context.Context, targets []string) ([]bitaxe.SystemInfo, error) {
	var data []bitaxe.SystemInfo

	for _, target := range targets {
		info, err := c.client.SystemInfo(ctx, target)
		if err != nil {
			return nil, fmt.Errorf("fetching system info from %s failed: %w", target, err)
		}

		data = append(data, info)
	}

	return data, nil
}
