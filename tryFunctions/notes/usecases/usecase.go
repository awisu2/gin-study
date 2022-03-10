package usecase

import (
	"context"

	"github.com/awisu2/gin-study/tryFunctions/domain"
)

func ProvideNoteUsecase() *NotesUsecase {
	return &NotesUsecase{}
}

type NotesUsecase struct {
}

func (usecase *NotesUsecase) GetById(ctx context.Context, id int64) (domain.Note, error) {
	return domain.Note{}, nil
}
