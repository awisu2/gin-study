package app

import (
	"github.com/awisu2/gin-study/tryFunctions/app/wire"
	_framework "github.com/awisu2/gin-study/tryFunctions/frameworks/gin"
)

func Run() {
	opts := &_framework.FrameworkGinOptions{
		Port: 8080,
	}
	opts.Fix()

	frameworks := wire.InitializeFrameworkGin(opts)
	wire.InitializeNotesGinAdapter(frameworks.Engine)

	frameworks.Run()
}
