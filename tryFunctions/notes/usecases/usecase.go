package usecase

import (
	"context"

	"github.com/awisu2/gin-study/tryFunctions/domain"
)

func ProvideNoteUsecase(repository domain.NoteRepository) *NotesUsecase {
	return &NotesUsecase{
		Repository: repository,
	}
}

type NotesUsecase struct {
	Repository domain.NoteRepository
}

func (usecase *NotesUsecase) GetById(ctx context.Context, id int64) (domain.Note, error) {
	return usecase.Repository.GetById(ctx, id)
}
