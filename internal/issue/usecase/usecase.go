package usecase

import (
	"github.com/sew810i9/opus/internal/domain"
)

type useCase struct {
	issueCommandRepository domain.IssueCommandRepository
	issueQueryRepository   domain.IssueQueryRepository
}

func NewUseCase(
	issueCommandRepository domain.IssueCommandRepository,
	issueQueryRepository domain.IssueQueryRepository,
) domain.IssueUseCase {
	return &useCase{
		issueCommandRepository: issueCommandRepository,
		issueQueryRepository:   issueQueryRepository,
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
