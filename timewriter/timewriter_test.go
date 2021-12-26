package timewriter

import (
	"github.com/zf1976/vlog"
	"testing"
)

func TestWrite(t *testing.T) {
	timeWriter := &TimeWriter{
		Dir:           "./log",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	log := vlog.NewLogger(timeWriter)
	log.Infof("hello vdns")
}
