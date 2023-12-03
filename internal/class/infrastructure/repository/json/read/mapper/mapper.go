package mapper

import (
	"github.com/rcgc/go-hexagonal.git/internal/class/domain"
	"github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/dto"
)

type Mapper struct{}

func (m Mapper) DTOClassToDomain(class dto.Class) domain.Class {
	return *domain.NewClass(
		class.ClassID,
		class.Title,
		class.CreationDate,
		class.Content,
		class.ReadTime,
	)
}