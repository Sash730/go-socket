package controller

import (
	"log"
	"net/http"
	"github.com/Sash730/go-socket/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/gorilla/websocket"
)

var (
	// Websocket http upgrader
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type RecentlyController struct {

}

func NewRecentlyController() *RecentlyController {
	return &RecentlyController{

	}
}

func (rc RecentlyController) ViewReport(w http.ResponseWriter, req *http.Request) {
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		//var msg model.Preference
		//err := json.Unmarshal([]byte(m), &msg)
		//
		//updateUserPreferences(msg)


		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func updateUserPreferences(p model.Preference) {
	session, err := mgo.Dial("optimus:optimus@mongo:27017/optimus_test")
	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("optimus_test").C("UserPreferences")

	result := model.UserPreferences{}
	err = c.Find(bson.M{"username": p.Username, "entity_id": p.EntityId}).One(&result)
	log.Println("result.Preferences before 0:", err, result)

	if err != nil {
		result = model.UserPreferences{}
		err = result.СreatePreferences(result, p)
	} else {
		err = model.UpdatePreferences(result, bson.ObjectIdHex(p.ReportId))
	}
	if err != nil {
		return
	}
}