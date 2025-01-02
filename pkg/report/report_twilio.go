package report

import (
	"context"
	"fmt"
	"sync"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
	"github.com/adamdecaf/bitaxe-stats/pkg/blockchain"
)

// TODO(adam): send a text when a new highest difficulty is found

func newTwilioReporter(conf TwilioConfig) (*twilioReporter, error) {
	return &twilioReporter{}, nil // TODO(adam):
}

type twilioReporter struct {
	// client // TODO(adam): twilio client

	mu                sync.RWMutex
	highestDifficulty blockchain.Difficulty
}

func (r *twilioReporter) SystemInfo(ctx context.Context, data []bitaxe.SystemInfo) error {

	for _, info := range data {
		newdiff, err := r.isHighestDifficulty(info)
		if err != nil {
			return fmt.Errorf("%s system info: %w", info.Hostname, err)
		}
		if newdiff != nil {
			// TODO(adam): send SMS
		}
	}

	return nil // TODO(adam):
}

func (r *twilioReporter) isHighestDifficulty(info bitaxe.SystemInfo) (*blockchain.Difficulty, error) {
	diff, err := blockchain.ParseDifficulty(info.BestDiff)
	if err != nil {
		return nil, fmt.Errorf("parsing difficulty: %w", err)
	}

	r.mu.RLock()

	if diff.RawValue > r.highestDifficulty.RawValue {
		r.mu.RUnlock()

		r.mu.Lock()
		r.highestDifficulty = diff
		r.mu.Unlock()

		return &diff, nil
	}

	// not a new difficulty benchmark
	r.mu.RUnlock()

	return nil, nil
}
