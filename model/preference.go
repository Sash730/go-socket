package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

type UserPreferences struct {
	ID              bson.ObjectId    `json:"id" bson:"_id"`
	User            *User            `json:"user" bson:"user"`
	RecentlyReports []RecentlyReport `json:"recently_reports" bson:"recently_reports"`
	OpenReports     []RecentlyReport `json:"open_reports" bson:"open_reports"`
}

type RecentlyReport struct {
	Slug string    `json:"user" bson:"user"`
	Type string    `json:"recently_reports" bson:"recently_reports"`
	Time time.Time `json:"when" bson:"when"`
}

//func (p *UserPreferences) СreatePreferences(e interface{}, pr Preference) error {
//	key := e.(UserPreferences)
//	key.ID = bson.NewObjectId()
//	key.Username = pr.Username
//	key.EntityId = pr.EntityId
//	key.When = bson.Now()
//	key.Preferences = createDbRef(bson.ObjectIdHex(pr.ReportId))
//	log.Println("key", key)
//	session, err := mgo.Dial("optimus:optimus@mongo:27017/optimus_test")
//	if err != nil {
//		log.Fatal("cannot dial mongo", err)
//	}
//	defer session.Close()
//	c := session.DB("optimus_test").C("UserPreferences")
//	_, err = c.UpsertId(key.ID, &key)
//	return error(err)
//}
//
//func UpdatePreferences(p UserPreferences, ReportId bson.ObjectId) error {
//	pr := p.Preferences
//	newPref := createDbRef(ReportId)
//
//	for i := 0; i < len(pr); i++ {
//		if pr[i].Id == ReportId {
//			pr = append(pr[:i], pr[i+1:]...)
//			break
//		}
//	}
//	pr = append(newPref, pr...)
//
//	session, err := mgo.Dial("optimus:optimus@mongo:27017/optimus_test")
//	if err != nil {
//		log.Fatal("cannot dial mongo", err)
//	}
//	defer session.Close()
//	c := session.DB("optimus_test").C("UserPreferences")
//	_, err = c.UpsertId(p.ID, &p)
//
//	return error(err)
//}
