package time_tracker

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type MongoRepository struct {
	collection *mgo.Collection
}

func (repository *MongoRepository) Insert(data interface{}) {
	repository.collection.Insert(bson.M{"name": data, "checkIn": time.Now()})
}
