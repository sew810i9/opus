package domain

type Issue struct {
	ID         int64
	UUID       string
	Name       string
	Project    *Project
	CreatedBy  *User
	AssignedTo *User
}

type IssueUseCase interface {
	Create() error
	Update() error
}

type IssueCommandRepository interface {
	Create()
	Update()
}

type IssueQueryRepository interface {
	GetByUUID()
}
