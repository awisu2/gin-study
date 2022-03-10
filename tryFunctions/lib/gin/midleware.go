package gin

import "github.com/gin-gonic/gin"

type MiddlewareOptions struct {
	Logger bool
	Recovery bool
}

func setMiddleware(engine *gin.Engine, opts *MiddlewareOptions) {
	// set Logger
	if opts.Logger {
		engine.Use(gin.Logger())
	}

	// set Recovery
	if opts.Recovery {
		engine.Use(gin.Recovery())
	}
}