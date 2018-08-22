package repository

import "github.com/globalsign/mgo"

type MongoPersonRepository struct {
	db *mgo.Database
}

func NewMongoPersonRepository(db *mgo.Database) *MongoPersonRepository {
	return &MongoPersonRepository{
		db: db,
	}
}
