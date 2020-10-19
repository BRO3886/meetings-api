package participant

import (
	"github.com/BRO3886/meetings-api/pkg/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository is an abstraction over lower level DB functions
type Repository interface {
	FindByID(id string) (*entities.Participant, error)
	Create(user *entities.Participant) (*entities.Participant, error)
	DoesEmailExist(email string) (bool, error)
	FindByEmail(email string) (*entities.Participant, error)
}

//NewRepo new repo
func NewRepo(coll *mongo.Collection) Repository {
	return &repo{
		Coll: coll,
	}
}
