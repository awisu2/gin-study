package app

import (
	"fmt"

	"github.com/awisu2/gin-study/tryFunctions/gin"
	"github.com/awisu2/goUtils/runtimes"
)

// application option
type AppOptions struct {
	Host       string
	Port       int
	Production bool
}

// fix option values
func (opt *AppOptions) Fix() {
	// at production
	if !opt.Production {
		if opt.Host == "" {
			info := runtimes.GetInfo()
			if info.OsInfo.IsWindows {
				opt.Host = "localhost"
			}
		}
	}
}

// get address
func (opt *AppOptions) Address() string {
	return fmt.Sprintf("%s:%d", opt.Host, opt.Port)
}

// run application
func Run(opts *AppOptions) {
	opts.Fix()
	engine := gin.CreateEngine()

	setRouter(engine)

	engine.Run(opts.Address())
}
