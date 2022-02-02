package domain

import "time"

type Project struct {
	ID          int64
	UUID        string
	Name        string
	Description string
	CreatedBy   *User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProjectUseCase interface {
	Create() error
	Update() error
}

type ProjectCommandRepository interface {
	Create()
	Update()
}

type ProjectQueryRepository interface {
	GetByUUID()
}
