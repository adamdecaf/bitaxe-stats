package report

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
	"github.com/adamdecaf/bitaxe-stats/pkg/blockchain"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func newTwilioReporter(conf TwilioConfig) (*twilioReporter, error) {
	if conf.AccountSid == "" || conf.AuthToken == "" {
		return nil, errors.New("missing AccountSid / AuthToken")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: conf.AccountSid,
		Password: conf.AuthToken,
	})

	return &twilioReporter{
		conf:   conf,
		client: client,
	}, nil
}

type twilioReporter struct {
	conf   TwilioConfig
	client *twilio.RestClient

	mu                sync.RWMutex
	highestDifficulty blockchain.Difficulty
}

func (r *twilioReporter) SystemInfo(ctx context.Context, data []bitaxe.SystemInfo) error {
	var found *bitaxe.SystemInfo

	// Have we observed before?
	previouslyEmpty := r.previouslyEmpty()

	for _, info := range data {
		newdiff, err := r.isHighestDifficulty(info)
		if err != nil {
			return fmt.Errorf("%s system info: %w", info.Hostname, err)
		}
		if newdiff != nil {
			found = &info
		}
	}

	// only alert on the highest of the best difficulties
	if !previouslyEmpty && found != nil {
		err := r.sendSMS(ctx, *found)
		if err != nil {
			return fmt.Errorf("alerting on %s newdiff from %s failed: %w", found.BestDiff, found.Hostname, err)
		}
	}

	r.logBestDifficulty()

	return nil
}

func (r *twilioReporter) previouslyEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.highestDifficulty.RawValue <= 1.0
}

func (r *twilioReporter) logBestDifficulty() {
	r.mu.RLock()
	defer r.mu.RUnlock()

	log.Printf("highest difficulty found: %v", r.highestDifficulty.String())
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

func (r *twilioReporter) sendSMS(ctx context.Context, info bitaxe.SystemInfo) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(r.conf.To)
	params.SetFrom(r.conf.From)
	params.SetBody(fmt.Sprintf("%s reached a new difficulty of %s", info.Hostname, info.BestDiff))

	_, err := r.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("sending SMS: %w", err)
	}
	return nil
}
