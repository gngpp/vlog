<p align="center">
	<a href="https://github.com/gngpp/vlog/releases/tag/v1.0.4"><img src="https://img.shields.io/github/v/release/gngpp/vlog?logo=github" alt="GitHub release" /></a>
</p>
<p align="center">
  <strong>Chinese</strong> | <a href="https://github.com/gngpp/vlog/blob/main/README.en.md">English</a>
</p>

# vlog

vlog使用Golang原生日志库封装和实现io.Writer支持每天滚动压缩日志文件

## 概述

* 基于实现 io.Writer。您可以轻松地在 golang log、GORM、grpclog 等中使用。
* 支持每日滚动日志，可以指定日志文件名的前缀，默认为进程名
* 旧文件压缩为gz格式

## 入门

**示例**

要使用 vlog，您可以：

```shell
go get -u github.com/gngpp/vlog
```

并像这样导入：

```go
package main

import (
	"github.com/gngpp/vlog"
	"github.com/gngpp/vlog/timewriter"
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

如果需要全局设置输出：

```go
package main

import (
	"github.com/gngpp/vlog"
	"github.com/gngpp/vlog/timewriter"
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

如果需要全局设置输出并同步到默认日志库：

```go
package main

import (
	"github.com/gngpp/vlog"
	"github.com/gngpp/vlog/timewriter"
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

## 参考

* [lumberject](https://github.com/natefinch/lumberjack)
