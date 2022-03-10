package gin

import (
	"net/http"
	"strconv"

	"github.com/awisu2/gin-study/tryFunctions/domain"
	"github.com/gin-gonic/gin"
)

func ProvideNotesGinAdapter(e *gin.Engine, usecase domain.NoteUsecase) *NotesGinAdapter {
	adapter := &NotesGinAdapter{
		usecase: usecase,
	}

	e.GET("/notes/:id", adapter.NotesGetById)

	return adapter
}

type NotesGinAdapter struct {
	engine  *gin.Engine
	usecase domain.NoteUsecase
}

func (a *NotesGinAdapter) NotesGetById(c *gin.Context) {
	ctx := c.Request.Context()
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {

	}

	id := int64(idP)

	note, err := a.usecase.GetById(ctx, id)
	if err != nil {

	}

	c.String(http.StatusOK, "Hello %v", note)
}
