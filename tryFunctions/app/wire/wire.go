//go:build wireinject

package wire

import (
	"github.com/awisu2/gin-study/tryFunctions/domain"
	_framework "github.com/awisu2/gin-study/tryFunctions/frameworks/gin"
	_notesAdapter "github.com/awisu2/gin-study/tryFunctions/notes/adapters/gin"
	_notesUsecase "github.com/awisu2/gin-study/tryFunctions/notes/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeFrameworkGin(opts *_framework.FrameworkGinOptions) *_framework.FrameworkGin {
	wire.Build(_framework.ProvideFrameworkGin)
	return &_framework.FrameworkGin{}
}

func InitializeNotesGinAdapter(e *gin.Engine) *_notesAdapter.NotesGinAdapter {
	wire.Build(
		_notesUsecase.ProvideNoteUsecase,
		_notesAdapter.ProvideNotesGinAdapter,
		wire.Bind(new(domain.NoteUsecase), new(*_notesUsecase.NotesUsecase)),
	)
	return &_notesAdapter.NotesGinAdapter{}
}
