package time_tracker

import (
	"labix.org/v2/mgo"
)

type MongoRepository struct {
	collection *mgo.Collection
}

func (repository *MongoRepository) Insert(data interface{}) {
	repository.collection.Insert(data)
}
