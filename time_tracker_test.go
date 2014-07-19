package time_tracker

import (
	/*"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"*/
	"testing"
)

type MockRepository struct {
	List  []interface{}
	Count int
}

func (m *MockRepository) Insert(data interface{}) {
	m.List[0] = data
	m.Count = 1
}

func TestItShouldCreateNewRecordWhenCheckInFirstTime(t *testing.T) {
	mockRepository := MockRepository{
		make([]interface{}, 1),
		0,
	}

	timeTracker := TimeTracker{&mockRepository}
	timeTracker.CheckIn("roofimon")

	roofimon_checkin := len(mockRepository.List)
	if roofimon_checkin != 1 {
		t.Errorf("Expect one record but get %v", roofimon_checkin)
	}
}

/*func xTestIntegrateMongoDBItShouldCreateNewRecordWhenCheckInFirstTime(t *testing.T) {
	test_session, _ := mgo.Dial("192.168.1.37")
	defer test_session.Close()
	db := test_session.DB("test_time_tracker")
	collection := db.C("dtac")
	defer collection.DropCollection()

	timeTracker := TimeTracker{&MongoRepository{collection}}
	timeTracker.CheckIn("roofimon")

	roofimon_checkin, _ := collection.Find(bson.M{"username": "roofimon"}).Count()
	if roofimon_checkin != 1 {
		t.Errorf("Expect one record but get %v", roofimon_checkin)
	}
}*/
