package report

import (
	"os"
	"strings"
)

type Config struct {
	Honeycomb *HoneycombConfig
	Twilio    *TwilioConfig
}

func DefaultConfig() Config {
	return Config{
		Honeycomb: defaultHoneycombConfig(),
		Twilio:    defaultTwilioConfig(),
	}
}

type HoneycombConfig struct {
	// TODO(adam):
}

func defaultHoneycombConfig() *HoneycombConfig {
	return nil // TODO(adam):
}

type TwilioConfig struct {
	AccountSid string
	AuthToken  string

	From string
	To   string
}

func defaultTwilioConfig() *TwilioConfig {
	var conf TwilioConfig

	conf.AccountSid = strings.TrimSpace(os.Getenv("TWILIO_ACCOUNT_SID"))
	conf.AuthToken = strings.TrimSpace(os.Getenv("TWILIO_AUTH_TOKEN"))

	conf.From = strings.TrimSpace(os.Getenv("TWILIO_FROM"))
	conf.To = strings.TrimSpace(os.Getenv("TWILIO_TO"))

	if conf.AccountSid != "" && conf.AuthToken != "" {
		return &conf
	}
	return nil
}
