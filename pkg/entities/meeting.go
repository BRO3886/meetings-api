package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Meeting struct
type Meeting struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title     string             `bson:"title" json:"title"`
	StartTime time.Time          `bson:"start_time" json:"start_time"`
	EndTime   time.Time          `bson:"end_time" json:"end_time"`
	Created   time.Time          `bson:"created_at" json:"created_at"`
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
	MeetingID     primitive.ObjectID `bson:"meeting_id,omitempty" json:"meeting_id"`
	PariticpantID primitive.ObjectID `bson:"p_id,omitempty" json:"p_id"`
	Response      UserResponse       `bson:"response" json:"response"`
}
