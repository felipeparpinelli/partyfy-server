package main

import (
	"net/http"
	"./controllers"
	"github.com/gorilla/mux"
)

type Spotify struct {
	controller controllers.SpotifyClient
}

type Party struct {
	controller controllers.PartyClient
}

func main() {

	spotify := &Spotify{controller: controllers.SpotifyClient{}}
	party := &Party{controller: controllers.PartyClient{}}

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/login", spotify.LoginHandler).Methods("GET")
	r.HandleFunc("/callback", spotify.CompleteAuth).Methods("GET")
	r.HandleFunc("/api/v1/users/{userID}/playlists", spotify.GetPlaylistsHandler).Methods("GET")
	r.HandleFunc("/api/v1/users/{userID}/playlists/{playlistID}", spotify.GetPlaylistHandler).Methods("GET")

	r.HandleFunc("/api/v1/collaborativeplaylist", party.GetCollaborativePlaylist).Methods("GET")
	r.HandleFunc("/api/v1/addtrack", party.AddTrack).Methods("POST", "OPTIONS")

	http.ListenAndServe(":8081", r)
}
