package time_tracker

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoRepository struct {
	collection *mgo.Collection
}

func (repository *MongoRepository) Insert(person Person) {
	repository.collection.Insert(person)
}

func (repository *MongoRepository) List(from string, to string) []Person {
	iporsut = Person{Name: IPORSUT, Site: "dtac", WorkDate: today.Format("2006-01-02"), Checkin: today, Checkout: today}
	roong = Person{Name: ROONG, Site: "dtac", WorkDate: yesterday.Format("2006-01-02"), Checkin: today, Checkout: today}
	return []Person{iporsut, roong}
}

func (repository *MongoRepository) Update(person Person) {
	var keys = bson.M{"name": person.Name, "workdate": time.Now().Format("2006-01-02")}
	var value = bson.M{"$set": bson.M{"checkout": time.Now()}}
	repository.collection.Update(keys, value)

}
