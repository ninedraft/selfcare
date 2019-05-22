package cli

import (
	"golang.org/x/net/proxy"
)

type Config struct {
	Config        string `json:"-" yaml:"-"`
	TodoistToken  string `json:"todoist-token" yaml:"todoist-token"`
	TelegramToken string `json:"telegram-token" yaml:"telegram-token"`
	DB            string `json:"db" yaml:"db"`
	proxy.Auth
	// flag definitions here
	// https://github.com/octago/sflags#flags-based-on-structures------
}
