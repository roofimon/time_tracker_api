package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Site     string        `bson:"site"`
	WorkDate string        `bson:"workdate"`
	Checkin  time.Time     `bson:"checkin"`
	Checkout time.Time     `bson:"checkout"`
}
