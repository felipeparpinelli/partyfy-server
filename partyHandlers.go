package main

import (
	"log"
	"encoding/json"
	"net/http"
	"./models"
)

func  (party *Party) GetCollaborativePlaylist(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	result, err := party.controller.GetCollaborativePlaylist()
	if err != nil {
		log.Fatal(err)
	}

	out, err := json.MarshalIndent(result, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(out)
}

func (party *Party) AddTrack(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	track := getTrack(r)

	err := party.controller.AddTrack(track)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTrack (r *http.Request) (models.Track) {

	track := models.Track{}
	track.UserID = r.FormValue("userID")
	track.TrackID = r.FormValue("trackID")
	track.TrackName = r.FormValue("trackName")

	return track
}