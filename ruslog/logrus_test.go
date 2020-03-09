package ruslog

import (
	"os"
	"testing"

	"github.com/json-iterator/go"
)

func TestPrintln(t *testing.T) {
	SetFormatter(&JSONFormatter{
		TimestampFormat:  "",
		DisableTimestamp: false,
		DataKey:          "",
		FieldMap:         nil,
		CallerPrettyfier: nil,
		PrettyPrint:      true,
	})

	SetReportCaller(true)

	var u = struct {
		Username string
		Sex      int
	}{
		Username: "linty",
		Sex:      1,
	}

	str, _ := jsoniter.MarshalToString(u)

	Println(str)

	logger := New()
	logger.SetFormatter(&TextFormatter{TimestampFormat: "2006/01/02 15:04:05"})
	logger.Println("ruslog.New()")

	var log = Logger{
		Out:   os.Stderr,
		Hooks: nil,
		Formatter: &TextFormatter{
			TimestampFormat: "2006/01/02 15:04:05",
		},
		ReportCaller: true,
		Level:        DebugLevel,
		ExitFunc:     nil,
	}

	NewEntry(&log).Println("ruslog.NewEntry")
}
