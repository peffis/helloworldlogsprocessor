package helloworldlogsprocessor

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"go.opentelemetry.io/collector/processor/processorhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "helloworldlogsprocessor"
	// The stability level of the processor.
	stability = component.StabilityLevelDevelopment
)

var (
	consumerCapabilities = consumer.Capabilities{MutatesData: true}
)

// NewFactory returns a new factory for the HelloWorld logs processor.
func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		processor.WithLogs(createLogsProcessor, stability))
}

func createDefaultConfig() component.Config {
	return &Config{
		BaseConfig: adapter.BaseConfig{
			Operators: []operator.Config{},
		},
	}
}

func createLogsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	cfg component.Config,
	nextConsumer consumer.Logs) (processor.Logs, error) {

	hwp := &helloWorldProcessor{set.Logger}

	return processorhelper.NewLogsProcessor(
		ctx,
		set,
		cfg,
		nextConsumer,
		hwp.processLogs,
		processorhelper.WithCapabilities(consumerCapabilities))
}
