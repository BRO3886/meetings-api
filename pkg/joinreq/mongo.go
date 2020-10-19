package joinreq

import (
	"context"

	"github.com/BRO3886/meetings-api/pkg"
	"github.com/BRO3886/meetings-api/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	Coll *mongo.Collection
}

func (r *repo) Create(req *entities.JoinRequest) (*entities.JoinRequest, error) {
	req.ID = primitive.NewObjectID()
	_, err := r.Coll.InsertOne(context.Background(), req)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return req, nil
}

func (r *repo) FindByID(id string) (*entities.JoinRequest, error) {
	req := &entities.JoinRequest{}
	mID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	err = r.Coll.FindOne(context.Background(), bson.M{"_id": mID}).Decode(req)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return req, nil
}
