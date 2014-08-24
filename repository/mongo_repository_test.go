package repository

import (
	"fmt"
	"testing"
	"time"
	"github.com/roofimon/time_tracker_api/model"

	. "github.com/iporsut/test_set"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const IPORSUT = "iporsut"
const ROONG = "roong"
const ROOF = "roof"

var (
	session         mgo.Session
	mongoRepository MongoRepository
	collection      *mgo.Collection
	today           = time.Now()
	yesterday       = today.Add(-24 * time.Hour)
	iporsut         = model.Person{Name: IPORSUT, Site: "dtac", WorkDate: today.Format(MongoDateFormat), Checkin: today, Checkout: today}
	roong           = model.Person{Name: ROONG, Site: "dtac", WorkDate: yesterday.Format(MongoDateFormat), Checkin: today, Checkout: today}
	roof            = model.Person{Name: ROOF, Site: "dtac", WorkDate: today.Format(MongoDateFormat), Checkin: today, Checkout: today}
)

func TestExampleSuite(t *testing.T) {
	RunSuite(S{}, t)
}

type S struct {
}

func (S) Before(t *testing.T) {
	setUp()
}

func (S) After(t *testing.T) {
	tearDown()
}

func (S) TestGetDailyReport(t *testing.T) {
	mongoRepository.Insert(iporsut)
	mongoRepository.Insert(roong)
	mongoRepository.Insert(roof)
	from := time.Now().Format(MongoDateFormat)
	to := time.Now().Format(MongoDateFormat)
	dailyReport := mongoRepository.List(from, to)
	if recordCount := len(dailyReport); recordCount != 2 {
		t.Errorf("Expect records to equal 2 but got %v", recordCount)
	}
}

func (S) TestInsertOneTimeTrackingRecordIntoMongo(t *testing.T) {
	//Act
	mongoRepository.Insert(iporsut)
	//Assert
	if NotOnlyOneRecordInCollection() {
		t.Error("Expect 1 but got something else")
	}
}

func (S) TestUpdateAnExistingData(t *testing.T) {
	mongoRepository.Insert(iporsut)
	time.Sleep(1 * time.Second)
	//Act
	mongoRepository.Update(iporsut)
	//Assert
	if CheckoutIsEqualToCurrentTime() {
		t.Errorf("Expect not to equal check in time %v  but got %v", iporsut.Checkin, iporsut.Checkout)
	}
}

func CheckoutIsEqualToCurrentTime() bool {
	err := collection.Find(bson.M{"name": IPORSUT, "workdate": time.Now().Format(MongoDateFormat)}).One(&iporsut)
	if err != nil {
		fmt.Print("Can't find any record match to keyword")
	}
	if iporsut.Checkout == iporsut.Checkin {
		return true
	} else {
		return false
	}
}

func NotOnlyOneRecordInCollection() bool {
	var result bool = false
	timeTrackingRecord, _ := collection.Find(bson.M{"name": IPORSUT}).Count()
	if timeTrackingRecord != 1 {
		result = true
	}
	return result
}

func setUp() {
	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	collection = session.DB("test_site").C("dtac")
	mongoRepository = MongoRepository{collection}
}

func tearDown() {
	session.Close()
	collection.DropCollection()
}
