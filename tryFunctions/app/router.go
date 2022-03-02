package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter(engine *gin.Engine) {

	engine.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// engine.GET("/someGet", getting)
	// engine.POST("/somePost", posting)
	// engine.PUT("/somePut", putting)
	// engine.DELETE("/someDelete", deleting)
	// engine.PATCH("/somePatch", patching)
	// engine.HEAD("/someHead", head)
	// engine.OPTIONS("/someOptions", options)

}
