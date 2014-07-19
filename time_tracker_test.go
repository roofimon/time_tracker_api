package time_tracker
import (
  "testing"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

func TestItShouldCreateNewRecordWhenCheckInFirstTime(t *testing.T) {
  test_session, _ := mgo.Dial("172.16.129.130")
  defer test_session.Close()
  db := test_session.DB("test_time_tracker")
  timeTracker := TimeTracker{db}
  timeTracker.CheckIn("roofimon")
  dtac_collection := db.C("dtac") 
  defer dtac_collection.DropCollection()
  roofimon_checkin, _ := dtac_collection.Find(bson.M{"username": "roofimon"}).Count()
  if roofimon_checkin != 1 {
    t.Errorf("Expect one record but get %v", roofimon_checkin)
  }
}

