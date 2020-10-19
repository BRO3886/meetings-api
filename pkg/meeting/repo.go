package meeting

import (
	"time"

	"github.com/BRO3886/meetings-api/pkg/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository for Meetings
type Repository interface {
	CreateMeeting(meeting *entities.Meeting) (*entities.Meeting, error)
	FindByID(id string) (*entities.Meeting, error)
	FindByParticipant(email string) (*[]entities.Meeting, error)
	FindByTimeStamp(startTime, endTime time.Time) (*[]entities.Meeting, error)
}

//NewRepo new repo
func NewRepo(coll *mongo.Collection) Repository {
	return &repo{
		Coll: coll,
	}
}
