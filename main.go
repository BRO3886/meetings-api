package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/meetings-api/api/handlers"

	"github.com/BRO3886/meetings-api/pkg/meeting"
	"github.com/BRO3886/meetings-api/utils"
	"github.com/joho/godotenv"
)

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 3000")
		return ":3000"
	}
	return ":" + port
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := utils.DBConnect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	fmt.Println("Connected to DB")

	//instantiate servemux
	r := http.NewServeMux()

	//get collections
	// pColl := db.Collection("participants")
	mColl := db.Collection("meetings")
	// rColl := db.Collection("joinRequests")

	//create repos and services
	mRepo := meeting.NewRepo(mColl)
	mSvc := meeting.NewService(mRepo)

	//mount routes
	handlers.MountMeetingRoutes(r, mSvc)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "all ok 🔥",
			"ping":    "pong",
		})
		return
	})

	log.Println("Started server")

	log.Fatal(http.ListenAndServe(getPort(), r))
}
