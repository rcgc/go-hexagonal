package mapper

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rcgc/go-hexagonal.git/internal/class/application/query"
	"github.com/rcgc/go-hexagonal.git/internal/class/domain"
	"github.com/rcgc/go-hexagonal.git/src/handler/class/view/contract"
)

type Mapper struct{}

func (m Mapper) DomainToResponse(classes domain.Class) contract.Response {

	return contract.Response{
		ClassID:      classes.ClassID(),
		Title:        classes.Title(),
		CreationDate: classes.CreationDate(),
		Content:      classes.Content(),
		ReadTime:     classes.ReadTime(),
	}
}

func (m Mapper) RequestToQuery(ctx *gin.Context) (query.View, error) {
	request := contract.Request{
		Email:   ctx.Param("email"),
		ClassID: ctx.Param("class_id"),
		Title:   ctx.Param("title"),
	}

	if strings.HasPrefix(request.Email, ":") ||
		strings.HasPrefix(request.ClassID, ":") ||
		strings.HasPrefix(request.Title, ":") {

		return query.View{}, errors.New("request empty")
	}

	return *query.NewView(
		request.Email,
		request.ClassID,
		request.Title,
	), nil
}