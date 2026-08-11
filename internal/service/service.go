package service

type Service struct {
	repository Repository
}

type Repository interface {
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
