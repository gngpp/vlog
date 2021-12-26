# vlog

vlog implements io.Writer to roll and compress log file time every day, and encapsulate the use of Golang native log library

## Overview

* implements io.Writer. You can easily use in golang log, GORM, grpclog etc.
* daily roll log, you can specific the log file name's prefix, default is process name
* compress to gz for old file

## Getting Started

**Example**

To use vlog, you can git clone [https://github.com/zf1976/vlog](https://github.com/zf1976/vlog), and import like this:

```go
package main

import (
	"vlog"
	"vlog/timewriter"
)

func main() {
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./log",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	log := vlog.NewLogger(timeWriter)
	log.Infof("hello vdns")
}
```
## Reference

* [lumberject](https://github.com/natefinch/lumberjack)
