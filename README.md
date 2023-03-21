# helloworld log processor

The purpose of this repo is to show how one can write an opentelemetry
collector log processor. It does not really do anything but to add
some "hello":"world" entry to all log records.

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
