package mysql

import "github.com/sew810i9/opus/internal/domain"

type queryRepository struct {
}

func NewQueryRepository() domain.ProjectQueryRepository {
	return &queryRepository{}
}

func (q *queryRepository) GetByUUID() {
	//TODO implement me
	panic("implement me")
}
