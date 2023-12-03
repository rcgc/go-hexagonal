package usecase

import (
	"github.com/rcgc/go-hexagonal.git/internal/class/application/command"
	"github.com/rcgc/go-hexagonal.git/internal/class/application/query"
	"github.com/rcgc/go-hexagonal.git/internal/class/domain"
)

type RepositoryViewClass interface {
	GetClassByClassID(classID string) (domain.Class, error)
}

type RepositoryUpdateClassesDone interface {
	UpdateClassesByEmail(cmd command.Update) error
}

type ViewUseCase struct {
	repositoryViewClass			RepositoryViewClass
	RepositoryUpdateClassesDone RepositoryUpdateClassesDone
}

func NewViewUseCase(repositoryViewClass RepositoryViewClass, RepositoryUpdateClassesDone RepositoryUpdateClassesDone) *ViewUseCase{
	return &ViewUseCase {
		repositoryViewClass: repositoryViewClass,
		RepositoryUpdateClassesDone: RepositoryUpdateClassesDone,
	}
}

func (uc ViewUseCase) Execute(qry query.View) (domain.Class, error) {
	domainClass, err := uc.repositoryViewClass.GetClassByClassID(qry.ClassID())
	if err != nil {
		return domain.Class{}, err
	}

	cmd := command.NewUpdate(
		qry.Email(),
		qry.ClassID(),
		qry.Title(),
	)
	err = uc.RepositoryUpdateClassesDone.UpdateClassesByEmail(*cmd)
	if err != nil {
		return domain.Class{}, err
	}

	return domainClass, nil
}