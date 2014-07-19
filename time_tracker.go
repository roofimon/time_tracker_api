package time_tracker
 import (
   "labix.org/v2/mgo"
   "labix.org/v2/mgo/bson"
 )

 type TimeTracker struct {
   db *mgo.Database
 }

 func (timeTracker TimeTracker) CheckIn(username string) {
  dtac_collection := timeTracker.db.C("dtac") 
  dtac_collection.Insert(bson.M{"username":username})
 }
