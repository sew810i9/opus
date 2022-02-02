package usecase

import (
	"github.com/sew810i9/opus/internal/domain"
)

type useCase struct {
	projectCommandRepository domain.ProjectCommandRepository
	projectQueryRepository   domain.ProjectQueryRepository
}

func NewUseCase(
	projectCommandRepository domain.ProjectCommandRepository,
	projectQueryRepository domain.ProjectQueryRepository,
) domain.ProjectUseCase {
	return &useCase{
		projectCommandRepository: projectCommandRepository,
		projectQueryRepository:   projectQueryRepository,
	}
}

func (u useCase) Create() error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) Update() error {
	//TODO implement me
	panic("implement me")
}
