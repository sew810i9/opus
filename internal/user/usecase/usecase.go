package usecase

import (
	"github.com/sew810i9/opus/internal/domain"
)

type useCase struct {
	userCommandRepository domain.UserCommandRepository
	userQueryRepository   domain.UserQueryRepository
}

func NewUseCase(
	userCommandRepository domain.UserCommandRepository,
	userQueryRepository domain.UserQueryRepository,
) domain.UserUseCase {
	return &useCase{
		userCommandRepository: userCommandRepository,
		userQueryRepository:   userQueryRepository,
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
