package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rcgc/go-hexagonal.git/internal/class/application/query"
	"github.com/rcgc/go-hexagonal.git/internal/class/domain"
	"github.com/rcgc/go-hexagonal.git/src/handler/class/view/contract"
)

type Mapper interface {
	RequestToQuery(ctx *gin.Context) (query.View, error)
	DomainToResponse(class domain.Class) contract.Response
}

type UseCase interface {
	Execute(qry query.View) (domain.Class, error)
}

type Handler struct {
	mapper  Mapper
	useCase UseCase
}

func NewHandler(mapper Mapper, useCase UseCase) *Handler {
	return &Handler{mapper: mapper, useCase: useCase}
}

func (h Handler) Handler(ginCTX *gin.Context) {
	qry, errBinding := h.mapper.RequestToQuery(ginCTX)
	if errBinding != nil {
		ginCTX.JSON(http.StatusBadRequest, nil)
		return
	}

	domainProfile, errorUseCase := h.useCase.Execute(qry)
	if errorUseCase != nil {
		ginCTX.JSON(http.StatusInternalServerError, domainProfile)
		return
	}

	response := h.mapper.DomainToResponse(domainProfile)
	ginCTX.JSON(http.StatusOK, response)
}