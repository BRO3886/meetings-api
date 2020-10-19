package meeting

import (
	"time"

	"github.com/BRO3886/meetings-api/pkg/entities"
)

//Service interface as an abstraction over lower level mongodb func
type Service interface {
	CreateMeeting(meeting *entities.Meeting) (*entities.Meeting, error)
	FindMeeting(id string) (*entities.Meeting, error)
	FindParticipantMeetings(email string) (*[]entities.Meeting, error)
	FindInRange(start, end time.Time) (*[]entities.Meeting, error)
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

func (s *service) CreateMeeting(meeting *entities.Meeting) (*entities.Meeting, error) {
	return s.repo.CreateMeeting(meeting)
}

func (s *service) FindMeeting(id string) (*entities.Meeting, error) {
	return s.repo.FindByID(id)
}

func (s *service) FindParticipantMeetings(email string) (*[]entities.Meeting, error) {
	return s.repo.FindByParticipant(email)
}

func (s *service) FindInRange(start, end time.Time) (*[]entities.Meeting, error) {
	return s.repo.FindByTimeStamp(start, end)
}
