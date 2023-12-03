package read

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rcgc/go-hexagonal.git/internal/class/domain"
	"github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/dto"
)

type Mapper interface {
	DTOClassToDomain(class dto.Class) domain.Class
}

type ClassRepositoryRead struct {
	mapper			Mapper
	filenameClasses string
}

func NewClassRepositoryRead(mapper Mapper, filenameClasses string) *ClassRepositoryRead {
	return &ClassRepositoryRead{
		mapper: mapper,
		filenameClasses: filenameClasses,
	}
}

func (r ClassRepositoryRead) GetClassByClassID(classID string) (domain.Class, error) {
	data, err := os.ReadFile(r.filenameClasses)
	if err != nil {
		return domain.Class{}, err
	}

	classes := make(map[string]dto.Class)
	err = json.Unmarshal(data, &classes)
	if err != nil {
		return domain.Class{}, err
	}

	classDTO, found := classes[classID]
	if !found{
		return domain.Class{}, fmt.Errorf("no class found for class_id: %s", classID)
	}

	classDone := r.mapper.DTOClassToDomain(classDTO)

	return classDone, nil
}