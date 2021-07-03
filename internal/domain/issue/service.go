package issue

type Service interface {
	Create() error
	Update() error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create() error {
	return nil
}

func (s *service) Update() error {
	return nil
}
