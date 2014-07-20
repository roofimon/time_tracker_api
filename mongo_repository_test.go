package time_tracker

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"testing"
    "time"
)

const IPORSUT = "iporsut"
var session mgo.Session
var mongoRepository MongoRepository
var collection *mgo.Collection

func setUp() {
	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("test_time_tracker").C("dtac")
	mongoRepository = MongoRepository{collection}
}


func tearDown() {
	session.Close()
	collection.DropCollection()
}

func TestInsertDataIntoMongo(t *testing.T) {
    setUp()
    defer tearDown()
	mongoRepository.Insert(IPORSUT)
	iporsutCheckin, _ := collection.Find(bson.M{"name": IPORSUT}).Count()
	if iporsutCheckin != 1 {
		t.Errorf("Expect 1 but got %v", iporsutCheckin)
	}
}

type Person struct {
    Checkout time.Time
}

func TestUpdateDateIntoMongo(t *testing.T) {
    setUp()
    defer tearDown()
	mongoRepository.Insert(IPORSUT)
    mongoRepository.Update(IPORSUT)
    var iporsut Person
    err := collection.Find(bson.M{"name": IPORSUT}).One(&iporsut)
    if err != nil {
        t.Error("Can't find any record match to keyword")
    }
    if iporsut.Checkout.Unix() == 0 {
        t.Errorf("Expect not to equal 0 but got %v", iporsut)
    }
}
