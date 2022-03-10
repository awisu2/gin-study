package app

import (
	"github.com/awisu2/gin-study/tryFunctions/app/wire"
	_framework "github.com/awisu2/gin-study/tryFunctions/frameworks/gin"
	repos "github.com/awisu2/gin-study/tryFunctions/notes/repositories"
)

func Run() {
	opts := &_framework.FrameworkGinOptions{
		Port: 8080,
	}
	opts.Fix()

	frameworks := wire.InitializeFrameworkGin(opts)
	notes := repos.MemoryNotes{
		1: {1, "a", "1 is a"},
		2: {2, "b", "2 is b"},
		3: {3, "c", "3 is c"},
	}
	wire.InitializeNotesGinAdapter(frameworks.Engine, notes)

	frameworks.Run()
}
