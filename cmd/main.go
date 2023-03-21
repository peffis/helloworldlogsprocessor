package main

import (
	"log"

	filelogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	"go.opentelemetry.io/collector/connector"
	forwardconnector "go.opentelemetry.io/collector/connector/forwardconnector"
	"go.opentelemetry.io/collector/exporter"
	loggingexporter "go.opentelemetry.io/collector/exporter/loggingexporter"
	otlpexporter "go.opentelemetry.io/collector/exporter/otlpexporter"
	otlphttpexporter "go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension"
	ballastextension "go.opentelemetry.io/collector/extension/ballastextension"
	zpagesextension "go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	memorylimiterprocessor "go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	"go.opentelemetry.io/collector/receiver"
	otlpreceiver "go.opentelemetry.io/collector/receiver/otlpreceiver"

	"go.opentelemetry.io/collector/component"

	"github.com/peffis/helloworldlogsprocessor"
)

func main() {
	factories, err := components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
	}

	info := component.BuildInfo{
		Command:     "otelcol-helloworldlog",
		Description: "OpenTelemetry Collector with helloworld log processor",
		Version:     "1.0.0",
	}

	if err := run(otelcol.CollectorSettings{BuildInfo: info, Factories: factories}); err != nil {
		log.Fatal(err)
	}
}

func run(params otelcol.CollectorSettings) error {
	return runInteractive(params)
}

func runInteractive(params otelcol.CollectorSettings) error {
	cmd := otelcol.NewCommand(params)
	if err := cmd.Execute(); err != nil {
		log.Fatalf("collector server run finished with error: %v", err)
	}

	return nil
}

func components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}

	factories.Extensions, err = extension.MakeFactoryMap(
		ballastextension.NewFactory(),
		zpagesextension.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Receivers, err = receiver.MakeFactoryMap(
		otlpreceiver.NewFactory(),
		filelogreceiver.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Exporters, err = exporter.MakeFactoryMap(
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Processors, err = processor.MakeFactoryMap(
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		helloworldlogsprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Connectors, err = connector.MakeFactoryMap(
		forwardconnector.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	return factories, nil
}
