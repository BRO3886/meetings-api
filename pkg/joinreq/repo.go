package joinreq

import (
	"github.com/BRO3886/meetings-api/pkg/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

//Repository for Join req
type Repository interface {
	Create(req *entities.JoinRequest) (*entities.JoinRequest, error)
	FindByID(id string) (*entities.JoinRequest, error)
}

//NewRepo new repo
func NewRepo(coll *mongo.Collection) Repository {
	return &repo{
		Coll: coll,
	}
}
