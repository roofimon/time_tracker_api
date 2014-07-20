package time_tracker

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"testing"
    "time"
)

func TestInsertDataIntoMongo(t *testing.T) {
	session, _ := mgo.Dial("localhost")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test_time_tracker").C("dtac")
	mongoRepository := MongoRepository{collection}
	mongoRepository.Insert("iporsut")
	defer collection.DropCollection()
	iporsutCheckin, _ := collection.Find(bson.M{"name": "iporsut"}).Count()
	if iporsutCheckin != 1 {
		t.Errorf("Expect 1 but got %v", iporsutCheckin)
	}
}

type Person struct {
    Checkout time.Time
}

func TestUpdateDateIntoMongo(t *testing.T) {
	session, _ := mgo.Dial("localhost")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test_time_tracker").C("dtac")
	mongoRepository := MongoRepository{collection}
	mongoRepository.Insert("iporsut")
    mongoRepository.Update("iporsut")
    defer collection.DropCollection()
    var iporsut Person
    err := collection.Find(bson.M{"name": "iporsut"}).One(&iporsut)
    if err != nil {
        t.Error("Can't find any record match to keyword")
    }
    if iporsut.Checkout.Unix() == 0 {
        t.Errorf("Expect not to equal 0 but got %v", iporsut)
    }
}
