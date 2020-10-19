package joinreq

import "github.com/BRO3886/meetings-api/pkg/entities"

//Service for join req
type Service interface {
	CreateRequest(req *entities.JoinRequest) (*entities.JoinRequest, error)
	FindByID(id string) (*entities.JoinRequest, error)
}

//NewService create new instance of service
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

type service struct {
	repo Repository
}

func (s *service) CreateRequest(req *entities.JoinRequest) (*entities.JoinRequest, error) {
	return s.repo.Create(req)
}

func (s *service) FindByID(id string) (*entities.JoinRequest, error) {
	return s.repo.FindByID(id)
}
