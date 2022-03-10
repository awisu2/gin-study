package usecase

import (
	"context"

	"github.com/awisu2/gin-study/tryFunctions/domain"
)

type MemoryNotes map[int64]domain.Note

func ProvideNotesRepository(datas MemoryNotes) *NotesMemoryRepository {
	return &NotesMemoryRepository{
		datas: datas,
	}
}

type NotesMemoryRepository struct {
	datas MemoryNotes
}

func (repository *NotesMemoryRepository) GetById(ctx context.Context, id int64) (domain.Note, error) {
	return repository.datas[id], nil
}
