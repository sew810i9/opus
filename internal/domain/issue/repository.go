package issue

import "github.com/jmoiron/sqlx"

type Repository interface {
	Create() error
	Update() error
}

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return &repository{}
}

func (r *repository) Create() error {
	return nil
}

func (r *repository) Update() error {
	return nil
}
