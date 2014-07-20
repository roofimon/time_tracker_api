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
var iporsut = Person{Name: IPORSUT, Site: "dtac", Checkin: time.Now().Unix(), Checkout: 0 }

type Person struct {
    Name string
    Site string
    Checkin int64
    Checkout int64 "checkout"
}

func setUp() { session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("test_time_tracker").C("dtac")
	mongoRepository = MongoRepository{collection}
}

func tearDown() {
	session.Close()
	collection.DropCollection()
}

func TestInsertOneTimeTrackingRecordIntoMongo(t *testing.T) {
    setUp()
    defer tearDown()

	mongoRepository.Insert(iporsut)

	if NotOnlyOneRecordInCollection() {
		t.Error("Expect 1 but got something else")
	}
}

func NotOnlyOneRecordInCollection() bool{
    var result bool = false
	timeTrackingRecord, _ := collection.Find(bson.M{"name": IPORSUT}).Count()
    if timeTrackingRecord != 1 {
      result =  true
    }
    return result
}

func TestUpdateAnExistungData(t *testing.T) {
    setUp()
    defer tearDown()
	mongoRepository.Insert(iporsut)

    mongoRepository.Update(IPORSUT)

    err := collection.Find(bson.M{"name": IPORSUT}).One(&iporsut)
    if err != nil {
        t.Error("Can't find any record match to keyword")
    }
    if iporsut.Checkout == 0 {
        t.Errorf("Expect to equal current date time (int64 format) but got %v", iporsut.Checkout)
    }
}
