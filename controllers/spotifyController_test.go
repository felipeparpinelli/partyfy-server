package controllers

import (
	"testing"
	"strings"
	"net/http"
	"fmt"
	"encoding/json"
	b64 "encoding/base64"
	"os"
	"github.com/zmb3/spotify"
)


const (
	SpotifyTokenURL = "https://accounts.spotify.com/api/token"
	SpotifyAPI = "https://api.spotify.com/v1/"
	SpotifyUsersURI = "users/"
	SpotifyPlaylistURI = "playlists/"
)

type ClientCredentials struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

var clientCrendentials = ClientCredentials{}
var playlists = spotify.SimplePlaylistPage{}
var playlist = spotify.FullPlaylist{}

/*
Não é possível testar usando a interface de spotifyController, pois seu fluxo de autenticação requer login
e senha do usuário. Para que o teste fosse possível, foi utilizado o client-credentials-flow. Os dados da struct
"ClientCredentials" é obtido a partir da requisição para "SpotifyTokenURL" com o base64 "SPOTIFY_ID" e "SPOTIFY_SECRET",
obtido das variáveis de ambiente. Com o "AccessToken", é possível obter os dados para os testes.
ref.: https://developer.spotify.com/web-api/authorization-guide/#client-credentials-flow
 */
func init () {
	clientCrendentials = getClientCredentials()
}

func TestPlaylists(t *testing.T) {

	getPlaylists("partyfyapp")

	if l := len(playlists.Playlists); l == 0 {
		t.Error("Didn't get any results")
		return
	}

	if playlists.Total != 5 {
		t.Error("Expected 5, got", playlists.Total)
	}

	p := playlists.Playlists[0]
	if p.Name != "guns n roses" {
		t.Error("Expected guns n roses, got", p.Name)
	}
	if p.Tracks.Total != 12 {
		t.Error("Expected 12 tracks, got", p.Tracks.Total)
	}
}

func TestPlaylist(t *testing.T) {

	getPlaylist("partyfyapp", "0UiPpsLEIyx5ezDimClUMi")

	if playlist.Name != "rock" {
		t.Error("Expected rock, got", playlist.Name)
	}

	if playlist.ID != "0UiPpsLEIyx5ezDimClUMi" {
		t.Error("Expected 0UiPpsLEIyx5ezDimClUMi, got", playlist.Name)
	}

	if playlist.Collaborative {
		t.Error("Expected false, got", playlist.Collaborative)
	}

	if playlist.Tracks.Total != 10 {
		t.Error("Expected 10 tracks, got", playlist.Tracks.Total)
	}

	expected := "Sultans Of Swing - Live At The BBC"
	actual := playlist.Tracks.Tracks[0].Track.Name
	if expected != actual {
		t.Errorf("Got '%s', expected '%s'\n", actual, expected)
	}
}

func spotifyBase64AuthCredential() string {
	clientID := os.Getenv("SPOTIFY_ID")
	clientSecret := os.Getenv("SPOTIFY_SECRET")
	data := clientID + ":" + clientSecret
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	return string(sEnc)
}

func getClientCredentials () ClientCredentials {

	base64 := spotifyBase64AuthCredential()

	body := strings.NewReader(`grant_type=client_credentials`)
	req, err := http.NewRequest("POST", SpotifyTokenURL, body)
	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Authorization", "Basic " + base64)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&clientCrendentials)
	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()
	return clientCrendentials

}

func getPlaylists (userID string) {

	req, err := http.NewRequest("GET", SpotifyAPI + SpotifyUsersURI + userID + "/" + SpotifyPlaylistURI, nil)
	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Authorization", "Bearer " + clientCrendentials.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&playlists)

	defer resp.Body.Close()
}

func getPlaylist (userID string, playlistID string) {

	req, err := http.NewRequest("GET", SpotifyAPI + SpotifyUsersURI + userID + "/" + SpotifyPlaylistURI + playlistID, nil)
	if err != nil {
		fmt.Print(err)
	}
	req.Header.Set("Authorization", "Bearer " + clientCrendentials.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&playlist)

	defer resp.Body.Close()
}
