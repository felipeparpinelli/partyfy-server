package config

import (
	"gopkg.in/mgo.v2"
	"log"
	"../models"
	"gopkg.in/mgo.v2/bson"
)

var DB *mgo.Database

var PartyCollection *mgo.Collection


func init() {
	session := getSession()
	DB = session.DB(DBName)
	PartyCollection = DB.C(DBCollectionName)
	//createCollection()
}

func getSession() *mgo.Session {
	session, err := mgo.Dial(DBURL)

	if err != nil {
		panic(err)
	}

	if err = session.Ping(); err != nil {
		panic(err)
	}

	return session
}

func createCollection() {
	err := PartyCollection.Insert(&models.Party{bson.NewObjectId(), DBPartyName, DBPartyOwner, []models.Track{}})
	if err != nil {
		log.Fatal(err)
	}
}