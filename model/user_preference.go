package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
	"github.com/globalsign/mgo"
	"log"
)

type Input struct {
	ReportId string `json:"report_id"`
	Token    string `json:"token"`
}

type Report struct {
	Collection string      `bson:"$ref"`
	Id         interface{} `bson:"$id"`
	Database   string      `bson:"$db,omitempty"`
}

type Preference struct {
	ReportId string `json:"report_id"`
	Username string `json:"username" bson:"username"`
	EntityId string `json:"entity_id" bson:"entity_id"`
}

type UserPreferences struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	Username    string        `json:"username" bson:"username"`
	EntityId    string        `json:"entity_id" bson:"entity_id"`
	When        time.Time     `json:"when" bson:"when"`
	Preferences []*mgo.DBRef  `json:"preferences" bson:"preferences"`
}

func getReport(reportId string) *Report {
	return &Report{
		Collection: "Report",
		Id: bson.ObjectId(reportId),
	}
}

func (p *UserPreferences) СreatePreferences (e interface{}, pr Preference) error {
	key := e.(UserPreferences)
	key.ID = bson.NewObjectId()
	key.Username = pr.Username
	key.EntityId = pr.EntityId
	key.When = bson.Now()
	key.Preferences = createDbRef(bson.ObjectIdHex(pr.ReportId))
	log.Println("key", key)
	session, err := mgo.Dial("optimus:optimus@mongo:27017/optimus_test")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer session.Close()
	c := session.DB("optimus_test").C("UserPreferences")
	_, err = c.UpsertId(key.ID, &key)
	return error(err)
}


func UpdatePreferences(p UserPreferences, ReportId bson.ObjectId) error {
	pr := p.Preferences
	newPref := createDbRef(ReportId)

	for i := 0; i < len(pr); i++ {
		if pr[i].Id == ReportId {
			pr = append(pr[:i], pr[i+1:]...)
			break
		}
	}
	pr = append(newPref, pr...)

	session, err := mgo.Dial("optimus:optimus@mongo:27017/optimus_test")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer session.Close()
	c := session.DB("optimus_test").C("UserPreferences")
	_, err = c.UpsertId(p.ID, &p)

	return error(err)
}

func createDbRef(ReportId bson.ObjectId) []*mgo.DBRef {
	reportRef := &mgo.DBRef{
		Collection: "Report",
		Id: ReportId,
	}

	return append([]*mgo.DBRef{reportRef})
}