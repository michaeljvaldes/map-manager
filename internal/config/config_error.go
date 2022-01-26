package config

import "fmt"

const defaultMessage = "error with parsing config"

type ConfigError struct {
	Err     error
	Context string
}

func (ce *ConfigError) Unwrap() error { return ce.Err }

func (ce *ConfigError) ErrorMessage() string {
	if len(ce.Context) > 0 {
		return fmt.Sprintf("%s: %s: %v", defaultMessage, ce.Context, ce.Err)
	} else {
		return fmt.Sprintf("%s: %v", defaultMessage, ce.Err)
	}
}
