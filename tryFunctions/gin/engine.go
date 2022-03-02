package gin

import (
	"github.com/gin-gonic/gin"
)

// CreateEngine() create *gin.Engine
func CreateEngine() *gin.Engine {
	return gin.Default()
}
