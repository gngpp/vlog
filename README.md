<p align="center">
	<a href="https://github.com/zf1976/vlog/blob/releases/latest"><img src="https://img.shields.io/github/v/release/ripperhe/Bob?logo=github" alt="GitHub release" /></a>
</p>
<p align="center">
  <strong>Chinese</strong> | <a href="https://github.com/zf1976/vlog/blob/main/README.en.md">English</a>
</p>

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
go get -u github.com/zf1976/vlog
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
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	log := vlog.New(timeWriter)
	log.Info("hello vlog")
}
```

If you need to set the output globally:

```go
package main

import (
	"github.com/zf1976/vlog"
	"github.com/zf1976/vlog/timewriter"
	"io"
	"os"
)

func main() {
	timeWriter := &timewriter.TimeWriter{
		Dir:           "./logs",
		Compress:      true,
		ReserveDay:    30,
		LogFilePrefix: "vlog",
	}
	w := io.MultiWriter(os.Stdout, timeWriter)
	// global settings
	vlog.SetOutput(w)
	logger := vlog.Default()
	logger.Info("hello vlog")
}
```

If you need to set the output globally and synchronize to the default log library:

```go
package main

import (
	"github.com/zf1976/vlog"
	"github.com/zf1976/vlog/timewriter"
	"io"
	"log"
	"os"
)

func main() {
	timeWriter := &timewriter.TimeWriter{
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
```

## Reference

* [lumberject](https://github.com/natefinch/lumberjack)
