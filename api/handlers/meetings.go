package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BRO3886/meetings-api/api/views"
	"github.com/BRO3886/meetings-api/pkg/entities"
	"github.com/BRO3886/meetings-api/pkg/meeting"
	"github.com/gorilla/mux"
)

func create(svc meeting.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
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
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		if r.Method != http.MethodGet {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
		idStr := r.URL.Path[len("/meetings/"):]
		m, err := svc.FindMeeting(idStr)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "meeting found",
			"meeting": m,
		})
		return
	}
}

func findByEmail(svc meeting.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		if r.Method != http.MethodGet {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
		email := r.URL.Query().Get("participant")
		meetings, err := svc.FindParticipantMeetings(email)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "meetings found",
			"meetings": meetings,
		})
		return
	}
}

func findInRange(svc meeting.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		if r.Method != http.MethodGet {
			views.Wrap(views.ErrMethodNotAllowed, w)
			return
		}
		q := r.URL.Query()
		start := q.Get("start")
		end := q.Get("end")

		layout := "2006-01-02T15:04:05.000Z"

		sTime, err := time.Parse(layout, start)
		if err != nil {
			views.Wrap(views.ErrBadRequst, w)
			return
		}

		eTime, err := time.Parse(layout, end)
		if err != nil {
			views.Wrap(views.ErrBadRequst, w)
			return
		}
		meetings, err := svc.FindInRange(sTime, eTime)
		if err != nil {
			views.Wrap(err, w)
			return
		}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "meetings found",
			"meetings": meetings,
		})
		return
	}
}

//MountMeetingRoutes to handle routes of meeting
func MountMeetingRoutes(r *mux.Router, svc meeting.Service) {
	//schedule meetings
	r.Handle("/meetings", create(svc))
	r.Handle("/meeting/{id}", find(svc))
	r.Handle("/meetings?participant={emailid}", findByEmail(svc))
	r.Handle("/meetings?start={starttime}&end={endtime}", findInRange(svc))
}
