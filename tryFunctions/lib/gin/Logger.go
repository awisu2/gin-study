package gin

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type LogOptions struct {
	LogFile string // log file path
	EnableColor bool
	Formatter func(gin.LogFormatterParams) string
}

// sample formatter
var SampleFormatter = func (param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func setLog(engine *gin.Engine, opt *LogOptions) {
	if !opt.EnableColor {
		gin.DisableConsoleColor()
	}
	if opt.LogFile != ""{
		f, _ := os.Create(opt.LogFile)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	if opt.Formatter != nil {
		engine.Use(gin.LoggerWithFormatter(opt.Formatter))
	}
}