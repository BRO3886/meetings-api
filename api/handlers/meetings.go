package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BRO3886/meetings-api/api/views"
	"github.com/BRO3886/meetings-api/pkg/entities"
	"github.com/BRO3886/meetings-api/pkg/meeting"
)

func create(svc meeting.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
		meeting := &entities.Meeting{}
		if err := json.NewDecoder(r.Body).Decode(meeting); err != nil {
			views.Wrap(err, w)
			return
		}
		meeting, err := svc.CreateMeeting(meeting)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "meeting created",
			"meeting": meeting,
		})
		return
	}
}

func find(svc meeting.Service) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
		
		return 
	}
}

//MountMeetingRoutes to handle routes of meeting
func MountMeetingRoutes(r *http.ServeMux, svc meeting.Service) {
	//schedule meetings
	r.Handle("/meetings", http.Handler(create(svc)))
}
