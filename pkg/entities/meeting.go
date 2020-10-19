package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Meeting struct
type Meeting struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	StartTime primitive.DateTime `bson:"start_time" json:"start_time"`
	EndTime   primitive.DateTime `bson:"end_time" json:"end_time"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}

//UserResponse enum
type UserResponse string

const (
	//Yes response
	Yes UserResponse = "yes"
	//No response
	No UserResponse = "no"
	//Maybe response
	Maybe UserResponse = "maybe"
)

//JoinRequest join-request for meetings
type JoinRequest struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MeetingID primitive.ObjectID `bson:"meeting_id,omitempty" json:"meeting_id"`
	EmailID   string             `bson:"p_id,omitempty" json:"p_id"`
	Response  UserResponse       `bson:"response" json:"response"`
}
