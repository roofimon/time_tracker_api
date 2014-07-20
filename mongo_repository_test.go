package time_tracker

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"testing"
)

func TestInsertDataIntoMongo(t *testing.T) {
	session, _ := mgo.Dial("192.168.2.40, 192.168.1.37")
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test_time_tracker").C("dtac")
	mongoRepository := MongoRepository{collection}
	mongoRepository.Insert("iporsut")
	defer collection.DropCollection()
	iporsut_checkin, _ := collection.Find(bson.M{"name": "iporsut"}).Count()
	if iporsut_checkin != 1 {
		t.Errorf("Expect 1 but got %v", iporsut_checkin)
	}
}
