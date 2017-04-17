package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"./config"
)

func (spotify *Spotify) GetPlaylistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	if vars["userID"] == "" {
		http.Error(w, "MISSING_USER_ID", http.StatusInternalServerError)
		return
	}

	playlists, err := spotify.controller.GetPlaylists(vars["userID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(playlists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (spotify *Spotify) GetPlaylistHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)

	if vars["userID"] == "" {
		http.Error(w, "MISSING_USER_ID", http.StatusInternalServerError)
		return
	}

	if vars["playlistID"] == "" {
		http.Error(w, "MISSING_PLAYLIST_ID", http.StatusInternalServerError)
		return
	}

	var playlist, err = spotify.controller.GetPlaylist(vars["userID"], vars["playlistID"])

	js, err := json.Marshal(playlist)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (spotify *Spotify) LoginHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	url := spotify.controller.Login(config.State)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (spotify *Spotify) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user := spotify.controller.CompleteAuth(w, r)

	defer http.Redirect(w, r, config.RedirectLoggedURL + user.ID, http.StatusSeeOther)
}
