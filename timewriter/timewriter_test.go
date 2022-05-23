package timewriter

import (
	"github.com/gngpp/vlog"
	"io"
	"log"
	"os"
	"testing"
)

func TestDefault(t *testing.T) {
	timeWriter := &TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	w := io.MultiWriter(os.Stdout, timeWriter)
	logger := vlog.New(w)
	logger.Info("hello vlog")
}

func TestGlobal(t *testing.T) {
	timeWriter := &TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	w := io.MultiWriter(os.Stdout, timeWriter)
	vlog.SetOutput(w)
	// global settings
	logger := vlog.Default()
	logger.Info("hello vlog")
}

func TestGlobalSync(t *testing.T) {
	timeWriter := &TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	w := io.MultiWriter(os.Stdout, timeWriter)
	// global settings
	vlog.SetSyncOutput(true)
	vlog.SetOutput(w)
	logger := vlog.Default()
	logger.Info("hello vlog")
	log.Println("hello vlog")
}
