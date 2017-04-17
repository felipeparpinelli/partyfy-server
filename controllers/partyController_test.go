package controllers

import (
	"testing"
	"../models"
)

var party = PartyClient{}

func TestGetCollaborativePlaylist(t *testing.T) {

	party, err := party.GetCollaborativePlaylist()
	if err != nil {
		t.Error("Error getting CollaborativePlaylist")
	}

	if party.Owner != "felipeparpinelli" {
		t.Error("Expected felipeparpinelli, got ", party.Owner)
	}

	if party.Name != "Festa" {
		t.Error("Expected Festa, got ", party.Name)
	}

	if party.Track[0].TrackName != "Sultans Of Swing - Live At The BBC" {
		t.Error("Expected Sultans Of Swing - Live At The BBC, got ", party.Track[0].TrackName)
	}

	if party.Track[0].UserID != "felipeparpinelli" {
		t.Error("Expected felipeparpinelli, got ", party.Track[0].UserID)
	}

	if party.Track[0].TrackID != "2Hlu3GtnPU847FZiXsxEW8" {
		t.Error("Expected 2Hlu3GtnPU847FZiXsxEW8, got ", party.Track[0].TrackID)
	}

}

func TestAddTrack(t *testing.T) {

	track := models.Track{}
	track.TrackID = "5CQ30WqJwcep0pYcV4AMNc"
	track.UserID = "felipeparpinelli"
	track.TrackName = "Stairway To Heaven"

	err := party.AddTrack(track)

	if err != nil {
		t.Error("Error inserting track")
	}
}
