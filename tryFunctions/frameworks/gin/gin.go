package gin

import (
	"fmt"

	appGin "github.com/awisu2/gin-study/tryFunctions/lib/gin"
	"github.com/awisu2/goUtils/runtimes"
	"github.com/gin-gonic/gin"
)

func ProvideFrameworkGin(opts *FrameworkGinOptions) *FrameworkGin {
	opts.Fix()
	engine := CreateEngine(opts)

	return &FrameworkGin{
		Engine:  engine,
		options: opts,
	}
}

type FrameworkGin struct {
	Engine  *gin.Engine
	options *FrameworkGinOptions
}

func (g *FrameworkGin) Run() error {
	return g.Engine.Run(g.options.Address())
}

// application option
type FrameworkGinOptions struct {
	Host       string
	Port       int
	Production bool
}

// fix option values
func (opts *FrameworkGinOptions) Fix() {
	// at no production
	if !opts.Production {
		if opts.Host == "" {
			info := runtimes.GetInfo()
			if info.OsInfo.IsWindows {
				opts.Host = "localhost"
			}
		}
	}
}

// get address
func (opts *FrameworkGinOptions) Address() string {
	return fmt.Sprintf("%s:%d", opts.Host, opts.Port)
}

func CreateEngine(opts *FrameworkGinOptions) *gin.Engine {
	return appGin.CreateEngine(&appGin.CreateEngineOptions{
		Middleware: &appGin.MiddlewareOptions{
			Logger:   true,
			Recovery: true,
		},
		Log: &appGin.LogOptions{
			LogFile:     "gin.log",
			EnableColor: false,
			Formatter:   appGin.SampleFormatter,
		},
	})
}
