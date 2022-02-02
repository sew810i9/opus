package domain

import "time"

type User struct {
	ID        int64
	UUID      string
	Name      string
	Email     string
	Password  string
	Role      int64
	CreatedAt time.Time
	LastLogin time.Time
	BlockedAt time.Time
	Status    int64 // blocked / online / offline
}

type UserUseCase interface {
	Create() error
	Update() error
}

type UserCommandRepository interface {
	Create()
	Update()
}

type UserQueryRepository interface {
	GetByUUID()
}
