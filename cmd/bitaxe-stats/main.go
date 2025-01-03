package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
	"github.com/adamdecaf/bitaxe-stats/pkg/collect"
	"github.com/adamdecaf/bitaxe-stats/pkg/report"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// Set up a channel to listen for system interrupt signals
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	client := bitaxe.NewClient(nil)
	collector := collect.NewCollector(client)

	targets := strings.Split(os.Getenv("BITAXE_TARGETS"), ",")
	if len(targets) == 1 && targets[0] == "" {
		log.Println("No targets provided. Set the BITAXE_TARGETS environment variable.")
		return
	}

	reportConf := report.DefaultConfig()
	reporter, err := report.NewReporters(reportConf)
	if err != nil {
		log.Fatalf("Failed to create reporter: %v", err)
	}

	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	if err := tick(ctx, targets, collector, reporter); err != nil {
		log.Printf("ERROR during first tick: %v\n", err)
	}

	log.Println("bitaxe-stats started. Press Ctrl+C to stop.")

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down gracefully...")
			return

		case <-ticker.C:
			// Perform periodic work
			if err := tick(ctx, targets, collector, reporter); err != nil {
				log.Printf("ERROR during tick: %v\n", err)
			}
		}
	}
}

func tick(ctx context.Context, targets []string, collector collect.Collector, reporter report.Reporter) error {
	data, err := collector.SystemInfo(ctx, targets)
	if err != nil {
		return fmt.Errorf("collecting data: %w", err)
	}

	if err := reporter.SystemInfo(ctx, data); err != nil {
		return fmt.Errorf("reporting system info: %w", err)
	}

	return nil
}
