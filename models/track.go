package models

type Track struct {
	TrackID string `bson:"trackid"`
	TrackName string `bson:"trackname"`
	UserID string `bson:"userid"`
}