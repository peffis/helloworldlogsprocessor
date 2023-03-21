package helloworldlogsprocessor

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
)

// Config defines configuration for Resource processor.
type Config struct {
	adapter.BaseConfig `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {
	return nil
}
