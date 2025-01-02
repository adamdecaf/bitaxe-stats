package report

import (
	"context"
	"fmt"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
)

type Reporter interface {
	SystemInfo(ctx context.Context, data []bitaxe.SystemInfo) error
}

func NewReporters(conf Config) (Reporter, error) {
	var reporters []Reporter

	switch {
	case conf.Honeycomb != nil:
		r, err := newHoneycombReporter(*conf.Honeycomb)
		if err != nil {
			return nil, fmt.Errorf("honeycomb: %w", err)
		}
		if r == nil {
			reporters = append(reporters, r)
		}

	case conf.Twilio != nil:
		r, err := newTwilioReporter(*conf.Twilio)
		if err != nil {
			return nil, fmt.Errorf("twilio: %w", err)
		}
		if r == nil {
			reporters = append(reporters, r)
		}
	}

	if len(reporters) == 0 {
		return nil, fmt.Errorf("no report Config found")
	}

	return &MultiReporter{
		reporters: reporters,
	}, nil
}
