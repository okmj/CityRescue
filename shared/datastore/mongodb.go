package datastore

import (
	"github.com/okeyonyia123/cityrescue/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//class name with fields
type MongoDBDatastore struct {
	*mgo.Session
}

//constructor
func NewMongoDBDatastore(url string) (*MongoDBDatastore, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDBDatastore{
		Session: session,
	}, nil
}

func (m *MongoDBDatastore) CreateUser(user *models.User) error {

	session := m.Copy()

	defer session.Close()
	userCollection := session.DB("cityrescue").C("User")
	err := userCollection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDBDatastore) GetUser(username string) (*models.User, error) {

	session := m.Copy()
	defer session.Close()
	userCollection := session.DB("cityrescue").C("User")
	u := models.User{}
	err := userCollection.Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil

}

func (m *MongoDBDatastore) Close() {
	m.Close()
}
