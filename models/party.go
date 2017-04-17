package models

import "gopkg.in/mgo.v2/bson"

type Party struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Name string `bson:"name"`
	Owner string `bson:"owner"`
	Track []Track `bson:"tracks"`
}