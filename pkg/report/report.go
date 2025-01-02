package report

import (
	"context"
	"fmt"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
)

type Reporter interface {
	SystemInfo(ctx context.Context, data []bitaxe.SystemInfo) error
}

func NewReporter(conf Config) (Reporter, error) {
	switch {
	case conf.Honeycomb != nil:
		return newHoneycombReporter(*conf.Honeycomb)
	}
	return nil, fmt.Errorf("no report Config found")
}
