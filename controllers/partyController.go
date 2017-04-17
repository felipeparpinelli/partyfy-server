package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"../models"
	"../config"
)


type Party interface {
	GetCollaborativePlaylist() (models.Party, error)
	AddTrack (track models.Track)
}

type PartyClient struct{}


/* Obtém a playlist da festa. Por enquanto, só é possível obter a playlist da festa "config.DBPartyName".
 No futuro, o usuário poderá escolher a festa, que será solicitado pelo client através de um parâmetro no GET */
func (c PartyClient) GetCollaborativePlaylist() (models.Party, error) {

	party := models.Party{}
	err := config.PartyCollection.Find(bson.M{"name": config.DBPartyName}).One(&party)

	return party, err
}

func (c PartyClient) AddTrack (track models.Track) (err error) {

	partyUpdate := bson.M{"name" : config.DBPartyName}
	PushTrack := bson.M{"$push": bson.M{"tracks": track}}
	err = config.PartyCollection.Update(partyUpdate, PushTrack)

	return err
}