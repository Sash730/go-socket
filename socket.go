package main

import (
	"log"
	"net/http"

	//"crypto/rsa"
	"encoding/json"
	//"fmt"
	//"github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/googollee/go-socket.io"
	//"github.com/mitchellh/mapstructure"
	//"io/ioutil"
	"time"
)

var (
	//verifyKey *rsa.PublicKey
)

//func init() {
//	verifyBytes, err := ioutil.ReadFile("ruir.pem")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

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

func (p *UserPreferences) createPreferences (e interface{}, pr Preference) error {
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

func createDbRef(ReportId bson.ObjectId) []*mgo.DBRef {
	reportRef := &mgo.DBRef{
		Collection: "Report",
		Id: ReportId,
	}

	return append([]*mgo.DBRef{reportRef})
}

func updatePreferences(p UserPreferences, ReportId bson.ObjectId) error {
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

func updateUserPreferences(p Preference) {
	session, err := mgo.Dial("optimus:optimus@mongo:27017/optimus_test")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("optimus_test").C("UserPreferences")

	result := UserPreferences{}
	err = c.Find(bson.M{"username": p.Username, "entity_id": p.EntityId}).One(&result)
	log.Println("result.Preferences before 0:", err, result)

	if err != nil {
		result = UserPreferences{}
		err = result.createPreferences(result, p)
	} else {
		err = updatePreferences(result, bson.ObjectIdHex(p.ReportId))
	}
	if err != nil {
		return
	}
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")
		//so.Join("activity")
		so.On("send report", func(m string) {
			var msg Preference
			err = json.Unmarshal([]byte(m), &msg)

			//token, err := jwt.Parse(msg.Token, func(token *jwt.Token) (interface{}, error) {
			//	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			//		return nil, fmt.Errorf("Unexpected signing method:    %v", token.Header["alg"])
			//	}
			//	return verifyKey, nil
			//})
			//if err != nil {
			//	log.Fatal(err)
			//}
			//log.Println("token:", token.Valid)
			////if token.Valid {
			////log.Println("user:", token)
			//var preferences UserPreferences
			//mapstructure.Decode(token.Claims, &preferences)

			updateUserPreferences(msg)

			log.Println("emit:", msg, err)
			//}

		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
