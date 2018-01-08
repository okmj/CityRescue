package models

import (
	"time"

	"github.com/okeyonyia123/cityrescue/shared/util"
)

type Post struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Username          string `json:"username" bson:"username"`
	Category          string `json:"category" bson:"category"`
	City              string `json:"city" bson:"city"`
	Address           string `json:"address" bson:"address"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModified" bson:"timestampModified"`
}

func NewPost(username string, category string, city string, address string) *Post {

	now := time.Now()
	unixTimestamp := now.Unix()
	post := Post{UUID: util.GenerateUUID(), Username: username, Category: category, City: city, Address: address, TimestampCreated: unixTimestamp}
	return &post
}
