package person

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"reflect"
	"testing"
	"time"
)

var (
	session    *mgo.Session
	collection *mgo.Collection
)

func setUpDB(t *testing.T) {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		t.Errorf("expect no error but was %v", err)
	}
	collection = session.DB("test_time_tracker").C("dtac")
}

func tearDownDB(t *testing.T) {
	err := collection.DropCollection()
	if err != nil {
		t.Errorf("expect no err but was %v", err)
	}
	session.Close()
}

func TestMongoRepositoryInsertPerson(t *testing.T) {
	setUpDB(t)
	defer session.Close()
	defer collection.DropCollection()

	var repository Repository = NewMongoRepository(collection)

	repository.Insert(Person{
		Name:     "iporsut",
		Site:     "dtac",
		CheckIn:  time.Now(),
		CheckOut: time.Time{},
	})

	personCheckIn, _ := collection.Find(bson.M{
		"name": "iporsut",
	}).Count()

	if personCheckIn != 1 {
		t.Errorf("Expect 1 but got %d", personCheckIn)
	}
}

func TestMongoRepositoryUpdatePerson(t *testing.T) {
	setUpDB(t)
	defer session.Close()
	defer collection.DropCollection()

	var p Person

	collection.Insert(&Person{
		Name:     "iporsut",
		Site:     "dtac",
		CheckIn:  time.Now(),
		CheckOut: time.Time{},
	})
	collection.Find(bson.M{"name": "iporsut"}).One(&p)

	checkOutTime := time.Now()
	p.CheckOut = checkOutTime

	var repository Repository = NewMongoRepository(collection)

	var iporsut, _ = repository.Update(p)

	if reflect.DeepEqual(iporsut.CheckOut, checkOutTime) {
		t.Errorf("Expect %v but got %v", checkOutTime, iporsut.CheckOut)
	}
}
