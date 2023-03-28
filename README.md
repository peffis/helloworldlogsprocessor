# helloworld log processor

The purpose of this repo is to show how one can write an opentelemetry
collector log processor. It does not really do anything but to add
some "hello":"world" entry to all log records.

The OpenTelemetry collector project has provided a tutorial for
writing an OTEL collector receiver for traces. Although the reason for
creating this helloworld log processor repo was to explore log
processor writing, the information in the receiver tutorial is very
helpful as it explains common concepts for components such as
life cycle and configuration. So before reading further here it is
suggested that you first read the OTEL receiver tutorial at
[https://opentelemetry.io/docs/collector/trace-receiver/](https://opentelemetry.io/docs/collector/trace-receiver/).


## building
```bash
go build cmd/main.go
```

## running
```bash
./main --config example.yaml
```
(this starts the collector with a file logger that reads from the file
"tmp.json" which will pass the data on to the helloworldlogsprocessor)

```bash
sh test.sh
```
(this writes some log entries to tmp.json in the form of structured
json log entries)
