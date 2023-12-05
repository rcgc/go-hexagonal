package dependence

import (
	useCaseClass "github.com/rcgc/go-hexagonal.git/internal/class/application/usecase"
	repositoryViewClass "github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/read"
	mapperRepositoryViewClass "github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/read/mapper"
	repositoryUpdateClass "github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/write"
	mapperRepositoryUpdateClass "github.com/rcgc/go-hexagonal.git/internal/class/infrastructure/repository/json/write/mapper"
	useCaseStudent "github.com/rcgc/go-hexagonal.git/internal/student/application/usecase"
	repositoryViewProfile "github.com/rcgc/go-hexagonal.git/internal/student/infrastructure/repository/json/read"
	mapperRepositoryViewProfile "github.com/rcgc/go-hexagonal.git/internal/student/infrastructure/repository/json/read/mapper"
	handlerViewClass "github.com/rcgc/go-hexagonal.git/src/handler/class/view"
	mapperViewClass "github.com/rcgc/go-hexagonal.git/src/handler/class/view/mapper"
	handlerViewProfile "github.com/rcgc/go-hexagonal.git/src/handler/student/view"
	mapperViewProfile "github.com/rcgc/go-hexagonal.git/src/handler/student/view/mapper"
)

type HandlerContainer struct {
	ViewProfileHandler handlerViewProfile.Handler
	ViewClassHandler   handlerViewClass.Handler
}

func NewWire() HandlerContainer {
	filenameProfile := "dbtest/StudentsProfile.json"
	filenameClasses := "dbtest/Classes.json"
	filenameClassesDone := "dbtest/StudentsClassesDone.json"

	repositoryClassRead := repositoryViewClass.NewClassRepositoryRead(
		mapperRepositoryViewClass.Mapper{},
		filenameClasses,
	)
	repositoryClassUpdate := repositoryUpdateClass.NewClassRepositoryWrite(
		mapperRepositoryUpdateClass.Mapper{},
		filenameClassesDone,
	)
	repositoryProfileRead := repositoryViewProfile.NewProfileRepositoryRead(
		mapperRepositoryViewProfile.Mapper{},
		filenameProfile,
		filenameClassesDone,
	)

	return HandlerContainer{
		ViewClassHandler: newWireViewClassHandler(
			*repositoryClassRead,
			*repositoryClassUpdate,
		),
		ViewProfileHandler: newWireViewProfileHandler(
			*repositoryProfileRead,
		),
	}
}

func newWireViewClassHandler(
	repositoryViewClass repositoryViewClass.ClassRepositoryRead,
	repositoryUpdateClass repositoryUpdateClass.ClassRepositoryWrite,
) handlerViewClass.Handler {

	useCaseView := useCaseClass.NewViewUseCase(
		repositoryViewClass,
		repositoryUpdateClass,
	)
	return *handlerViewClass.NewHandler(
		mapperViewClass.Mapper{},
		useCaseView,
	)
}

func newWireViewProfileHandler(
	repositoryViewProfile repositoryViewProfile.ProfileRepositoryRead,
) handlerViewProfile.Handler {
	useCaseViewProfile := useCaseStudent.NewViewUseCase(
		repositoryViewProfile,
	)

	return *handlerViewProfile.NewHandler(
		mapperViewProfile.Mapper{},
		useCaseViewProfile,
	)
}