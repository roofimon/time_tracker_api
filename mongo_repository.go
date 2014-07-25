package time_tracker

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type MongoRepository struct {
	collection *mgo.Collection
}

func (repository *MongoRepository) Insert(person Person) {
	repository.collection.Insert(person)
}

func (repository *MongoRepository) Update(person Person) {
	var keys = bson.M{"name": person.Name, "workdate": time.Now().Format("2006-01-02")}
	var value = bson.M{"$set": bson.M{"checkout": time.Now()}}
	repository.collection.Update(keys, value)
}
