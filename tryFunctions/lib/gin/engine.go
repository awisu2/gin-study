package gin

import (
	"github.com/gin-gonic/gin"
)

type CreateEngineOptions struct {
	Middleware *MiddlewareOptions
	Log *LogOptions
}


// CreateEngine() create *gin.Engine
func CreateEngine(opts *CreateEngineOptions) *gin.Engine {
	engine := gin.New()

	if opts.Middleware != nil{
		setMiddleware(engine, opts.Middleware)
	}

	if opts.Log != nil {
		setLog(engine, opts.Log)
	}

	return engine
}
