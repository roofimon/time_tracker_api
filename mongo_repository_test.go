package time_tracker

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"testing"
    "time"
    "fmt"
)

const IPORSUT = "iporsut"
var session mgo.Session
var mongoRepository MongoRepository
var collection *mgo.Collection
var iporsut = Person{Name: IPORSUT, Site: "dtac", Checkin: time.Now().Unix(), Checkout: 0 }

func TestInsertOneTimeTrackingRecordIntoMongo(t *testing.T) {
    //Arrange
    setUp()
    defer tearDown()
    //Act
	mongoRepository.Insert(iporsut)
    //Assert
	if NotOnlyOneRecordInCollection() {
		t.Error("Expect 1 but got something else")
	}
}

func TestUpdateAnExistungData(t *testing.T) {
    //Arrange
    setUp()
    defer tearDown()
	mongoRepository.Insert(iporsut)
    //Act
    mongoRepository.Update(IPORSUT)
    //Assert
    if CheckoutIsNotEqualToCurrentTime() {
        t.Errorf("Expect to equal current date time (int64 format) but got %v", iporsut.Checkout)
    }
}

func CheckoutIsNotEqualToCurrentTime() bool {
    err := collection.Find(bson.M{"name": IPORSUT}).One(&iporsut)
    if err != nil {
        fmt.Print("Can't find any record match to keyword")
    }
    if iporsut.Checkout == 0 {
      return true
    }else{
      return false
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

func setUp() { session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("test_time_tracker").C("dtac")
	mongoRepository = MongoRepository{collection}
}

func tearDown() {
	session.Close()
	collection.DropCollection()
}
