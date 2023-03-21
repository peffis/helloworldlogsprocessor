package helloworldlogsprocessor

import (
	"context"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
	"log"
)

type helloWorldProcessor struct {
	logger *zap.Logger
}

func (hwp *helloWorldProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	// we print log records and add the field "hello" with the value "world"
	log.Printf("hello world: received %d log records", ld.LogRecordCount())

	rls := ld.ResourceLogs()
	for i := 0; i < rls.Len(); i++ {
		rs := rls.At(i)
		ilss := rs.ScopeLogs()
		for j := 0; j < ilss.Len(); j++ {
			ils := ilss.At(j)
			logs := ils.LogRecords()

			for k := 0; k < logs.Len(); k++ {
				log.Printf("log record: %d", k)
				lr := logs.At(k)
				attrs := lr.Attributes()
				attrs.PutStr("hello", "world")
				attrs.Range(func(k string, v pcommon.Value) bool {
					log.Printf("   attr: %s = %v", k, v)
					return true
				})
			}
		}
	}
	return ld, nil
}
