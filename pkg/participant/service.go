package participant

import (
	"github.com/BRO3886/meetings-api/pkg"
	"github.com/BRO3886/meetings-api/utils"

	"github.com/BRO3886/meetings-api/pkg/entities"
)

//Service is as abstarction over repo
type Service interface {
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

//Validate email
func validate(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return utils.EmailRegex.MatchString(email)
}

func (s *service) Register(user *entities.Participant) (*entities.Participant, error) {
	validate := validate(user.Email)
	if !validate {
		return nil, pkg.ErrEmail
	}
	exists, err := s.repo.DoesEmailExist(user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, pkg.ErrExists
	}
	return s.repo.Create(user)
}

func (s *service) GetPartipantByID(id string) (*entities.Participant, error) {
	return s.repo.FindByID(id)
}

func (s *service) GetPartipantByEmail(email string) (*entities.Participant, error) {
	return s.repo.FindByEmail(email)
}
