package datastore

import (
	"errors"

	"github.com/okeyonyia123/cityrescue/models"
)

type Datastore interface {
	CreateUser(user *models.User) error
	CreateHelper(user *models.User) error
	GetUser(username string) (*models.User, error)
	GetHelper(username string) (*models.User, error)
	CreatePost(post *models.Post) error
	GetPost(username string) (*[]models.Post, error)
	GetAllPost() (*[]models.Post, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MYSQL:
		return NewMySQLDS(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	//case REDIS:
	//return NewRedisDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist!")
	}

}
