package test

import (
	"github.com/zf1976/vlog"
	"testing"
)

func TestLog(t *testing.T) {
	logger := vlog.Default()
	logger.Infof("info")
	logger.Infof("info value:%s", "test")

}

func TestSetLevel(t *testing.T) {
	vlog.SetLevel("trace")
}

func TestTrace(t *testing.T) {
	logger := vlog.Default()
	vlog.SetLevel("trace")
	logger.Trace("trace")
	logger.SetLevel("off")
	logger.Trace("trace")
}

func TestTracef(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("trace")
	logger.Tracef("tracef")
	logger.SetLevel("off")
	if logger.IsTraceEnabled() {
		logger.Tracef("tracef")
	}
}

func TestDebug(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("debug")
	logger.Debug("debug")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestDebugf(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("debug")
	logger.Debugf("debugf")
	logger.SetLevel("off")
	logger.Debug("debug")
}

func TestInfo(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("info")
	logger.Info("info")
	logger.SetLevel("off")
	logger.Info("info")
}

func TestInfof(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("info")
	logger.Infof("infof")
	logger.SetLevel("off")
	logger.Infof("infof")
}

func TestWarn(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("warn")
	logger.Warn("warn")
	logger.SetLevel("off")
	logger.Warn("warn")
}

func TestWarnf(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("warn")
	logger.Warnf("warnf")
	logger.SetLevel("off")
	logger.Warnf("warnf")
}

func TestError(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("err")
	logger.Error("err")
	logger.SetLevel("off")
	logger.Error("err")
}

func TestErrorf(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel(vlog.Level.OFF)
	logger.Errorf("errorf")
	logger.SetLevel("off")
	logger.Errorf("errorf")
}

func TestFatal(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel(vlog.Level.OFF)
	logger.Fatal("fatal")
	logger.SetLevel("off")
	logger.Fatal("fatal")
}

func TestFatalf(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel(vlog.Level.OFF)
	logger.Fatalf("fatalf")
	logger.SetLevel("off")
	logger.Errorf("fatalf")
}

func TestIsTraceEnabled(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("trace")

	if !logger.IsTraceEnabled() {
		t.FailNow()

		return
	}
}

func TestIsDebugEnabled(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("debug")

	if !logger.IsDebugEnabled() {
		t.FailNow()

		return
	}
}

func TestIsWarnEnabled(t *testing.T) {
	logger := vlog.Default()
	logger.SetLevel("warn")
	if !logger.IsWarnEnabled() {
		t.FailNow()
		return
	}
}
