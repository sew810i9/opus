package mysql

import "github.com/sew810i9/opus/internal/domain"

type commandRepository struct {
}

func NewCommandRepository() domain.IssueCommandRepository {
	return &commandRepository{}
}

func (c *commandRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (c *commandRepository) Update() {
	//TODO implement me
	panic("implement me")
}
