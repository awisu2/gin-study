//go:build wireinject

package wire

import (
	"github.com/awisu2/gin-study/tryFunctions/domain"
	_framework "github.com/awisu2/gin-study/tryFunctions/frameworks/gin"
	_notesGinAdapter "github.com/awisu2/gin-study/tryFunctions/notes/adapters/gin"
	_notesRepository "github.com/awisu2/gin-study/tryFunctions/notes/repositories"
	_notesUsecase "github.com/awisu2/gin-study/tryFunctions/notes/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeFrameworkGin(opts *_framework.FrameworkGinOptions) *_framework.FrameworkGin {
	wire.Build(_framework.ProvideFrameworkGin)

	return &_framework.FrameworkGin{}
}

func InitializeNotesGinAdapter(e *gin.Engine, datas _notesRepository.MemoryNotes) *_notesGinAdapter.NotesGinAdapter {
	wire.Build(
		_notesRepository.ProvideNotesRepository,
		_notesUsecase.ProvideNoteUsecase,
		_notesGinAdapter.ProvideNotesGinAdapter,
		wire.Bind(new(domain.NoteRepository), new(*_notesRepository.NotesMemoryRepository)),
		wire.Bind(new(domain.NoteUsecase), new(*_notesUsecase.NotesUsecase)),
	)
	return &_notesGinAdapter.NotesGinAdapter{}
}
