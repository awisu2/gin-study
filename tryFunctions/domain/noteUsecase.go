package domain

import "context"

type NoteUsecase interface {
	GetById(ctx context.Context, id int64) (Note, error)
}

type NoteRepository interface {
	GetById(ctx context.Context, id int64) (Note, error)
}
