package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

//Participant struct
type Participant struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Email string             `bson:"email" json:"email"`
}
