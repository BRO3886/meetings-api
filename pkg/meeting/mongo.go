package meeting

import (
	"context"
	"time"

	"github.com/BRO3886/meetings-api/pkg"
	"github.com/BRO3886/meetings-api/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	Coll *mongo.Collection
}

func (r *repo) CreateMeeting(meeting *entities.Meeting) (*entities.Meeting, error) {
	meeting.ID = primitive.NewObjectID()
	meeting.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err := r.Coll.InsertOne(context.Background(), meeting)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return meeting, nil
}

func (r *repo) FindByID(id string) (*entities.Meeting, error) {
	meeting := &entities.Meeting{}
	mID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	err = r.Coll.FindOne(context.Background(), bson.M{"_id": mID}).Decode(meeting)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return meeting, nil
}

func (r *repo) FindByParticipant(email string) (*[]entities.Meeting, error) {
	usr := &entities.Participant{}
	usr.Email = email
	err := r.Coll.FindOne(context.Background(), bson.M{"email": email}).Decode(usr)
	if err != nil {
		return nil, pkg.ErrDoesNotExist
	}
	cursor, err := r.Coll.Find(context.Background(), bson.M{"p_id": email})
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	var meetings []entities.Meeting
	for cursor.Next(context.TODO()) {
		var req entities.JoinRequest
		_ = cursor.Decode(req)
		meeting, err := r.FindByID(req.MeetingID.String())
		if err != nil {
			return nil, pkg.ErrDatabase
		}
		meetings = append(meetings, *meeting)
	}
	return &meetings, nil
}

func (r *repo) FindByTimeStamp(startTime, endTime time.Time) (*[]entities.Meeting, error) {
	//TODO
	var meetings []entities.Meeting
	return &meetings, nil
}
