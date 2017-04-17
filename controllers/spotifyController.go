package controllers

import (
	"github.com/zmb3/spotify"
	"net/http"
	"log"
	"../config"
	"fmt"
)

type Spotify interface {
	GetPlaylists(string) (*spotify.SimplePlaylistPage, error)
	GetPlaylist(string, string) (*spotify.FullPlaylist, error)
	Login(string) (string)
	CompleteAuth (http.ResponseWriter, http.Request) (*spotify.PrivateUser)
}

type SpotifyClient struct{}

type SpotifyController struct {
	client *spotify.Client
}

var (
	auth = spotify.NewAuthenticator(config.RedirectURI, spotify.ScopePlaylistReadPrivate)
	spotifyController = new(SpotifyController)
)

func (c SpotifyClient) Login (state string) (string) {

	url := auth.AuthURL(state)
	return url
}

func newSpotifyClient(client *spotify.Client) *SpotifyController {
	return &SpotifyController{client}
}

func (c SpotifyClient) CompleteAuth (w http.ResponseWriter, r *http.Request) (*spotify.PrivateUser) {

	token, err := auth.Token(config.State, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != config.State{
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, config.State)
	}

	client := auth.NewClient(token)
	spotifyController = newSpotifyClient(&client)
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	return user
}

func (c SpotifyClient) GetPlaylist(userID string, playlistID string) (*spotify.FullPlaylist, error) {

	var playlist, err = spotifyController.client.GetPlaylist(userID, playlistID)

	if err != nil {
		return nil, fmt.Errorf("playlists: spotify api error, %s", err)
	}

	return playlist, nil

}

func (c SpotifyClient) GetPlaylists(user string) (*spotify.SimplePlaylistPage, error) {

	var playlists, err = spotifyController.client.GetPlaylistsForUser(user)

	if err != nil {
		return nil, fmt.Errorf("playlists: spotify api error, %s", err)
	}

	return playlists, nil
}