package main

import (
	"context"

	"github.com/adamdecaf/bitaxe-stats/pkg/bitaxe"
	"github.com/adamdecaf/bitaxe-stats/pkg/collect"
	"github.com/adamdecaf/bitaxe-stats/pkg/report"
)

func main() {
	ctx := context.Background()

	var baseAddress string // TODO(adam):
	client := bitaxe.NewClient(baseAddress, nil)

	collector := collect.NewCollector(client)

	var targets []string
	data, err := collector.SystemInfo(ctx, targets)
	if err != nil {
		panic(err)
	}

	var reportConf report.Config
	reporter, err := report.NewReporter(reportConf)
	if err != nil {
		panic(err)
	}

	err = reporter.SystemInfo(ctx, data)
	if err != nil {
		panic(err)
	}
}
