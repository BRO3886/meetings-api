package participant

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

func (r *repo) FindByID(id string) (*entities.Participant, error) {
	user := &entities.Participant{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.Coll.FindOne(context.Background(), bson.M{"_id": objID}).Decode(user)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) Create(user *entities.Participant) (*entities.Participant, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Coll.InsertOne(context.Background(), user)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) DoesEmailExist(email string) (bool, error) {
	count, err := r.Coll.CountDocuments(context.Background(), bson.M{"email": email})
	if err != nil {
		return true, pkg.ErrDatabase
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *repo) FindByEmail(email string) (*entities.Participant, error) {
	user := &entities.Participant{}
	err := r.Coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}
