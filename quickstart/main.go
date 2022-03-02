package main

import (
	"github.com/awisu2/goUtils/runtimes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	addr := ":8080"
	if runtimes.GetInfo().OsInfo.IsWindows {
		addr = "localhost:8080" // avoid windows security alert
	}

	r.Run(addr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
