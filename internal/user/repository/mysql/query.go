package mysql

import "github.com/sew810i9/opus/internal/domain"

type QueryRepository struct {
}

func NewQueryRepository() domain.UserQueryRepository {
	return &QueryRepository{}
}

func (q QueryRepository) GetByUUID() {
	//TODO implement me
	panic("implement me")
}
