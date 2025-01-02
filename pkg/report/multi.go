package report

import (
	"context"
	"fmt"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
)

type MultiReporter struct {
	reporters []Reporter
}

func (m *MultiReporter) SystemInfo(ctx context.Context, data []bitaxe.SystemInfo) error {
	for idx := range m.reporters {
		err := m.reporters[idx].SystemInfo(ctx, data)
		if err != nil {
			return fmt.Errorf("multireporter[%d]: %w", idx, err)
		}
	}
	return nil
}
