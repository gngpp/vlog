# vlog

vlog implements io.Writer to roll and compress log file time every day, and encapsulate the use of Golang native log library

## Overview

* implements io.Writer. You can easily use in golang log, GORM, grpclog etc.
* daily roll log, you can specific the log file name's prefix, default is process name
* compress to gz for old file

## Getting Started

**Example**

To use vlog, you can:

```shell
go get github.com/zf1976/vlog
```

and import like this:

```go
package main

import (
	"github.com/zf1976/vlog"
	"github.com/zf1976/vlog/timewriter"
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
