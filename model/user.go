package model

type User struct {
	Username   string `json:"username" bson:"username"`
	EntityID   string `json:"entity_id" bson:"entity_id"`
	EntityType string `json:"entity_type" bson:"entity_type"`
	UserID     string `json:"user_id" bson:"user_id"`
	UserType   string `json:"user_type" bson:"user_type"`
	NID        string `json:"nid" bson:"nid"`
	Roles      []string
}
