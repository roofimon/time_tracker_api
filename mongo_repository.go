package time_tracker

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoRepository struct {
	collection *mgo.Collection
}

func (repository *MongoRepository) Insert(person Person) {
	repository.collection.Insert(person)
}

func (repository *MongoRepository) List(from string, to string) []Person {
	var result []Person
	_ = repository.collection.Find(bson.M{"workdate": from}).All(&result)
	return result
}

func (repository *MongoRepository) Update(person Person) {
	var keys = bson.M{"name": person.Name, "workdate": time.Now().Format("2006-01-02")}
	var value = bson.M{"$set": bson.M{"checkout": time.Now()}}
	repository.collection.Update(keys, value)

}
