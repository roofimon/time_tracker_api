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
	repository.collection.Insert(data)
}

func (repository *MongoRepository) Update(data interface{}) {
    repository.collection.Update(bson.M{"name": data}, bson.M{"$set":bson.M{"checkout": time.Now().Unix()}})
}
