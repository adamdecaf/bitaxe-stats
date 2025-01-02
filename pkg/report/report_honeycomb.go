package report

import (
	"context"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
)

func newHoneycombReporter(conf HoneycombConfig) (*honeycombReporter, error) {
	return &honeycombReporter{}, nil // TODO(adam):
}

type honeycombReporter struct {
	// TODO(adam):
}

func (h *honeycombReporter) SystemInfo(ctx context.Context, data []bitaxe.SystemInfo) error {
	return nil // TODO(adam):
}
