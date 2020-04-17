package couchdb

import (
	"context"
	"fmt"

	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v3"
	"github.com/spf13/viper"
	"gitlab.com/renodesper/efishery-cli-app/models"
)

type (
	couchdb struct {
		db *kivik.DB
	}

	// CouchDBRepository ...
	CouchDBRepository interface {
		GetTasks() (*kivik.Rows, error)
		AddTask(task *models.Task) error
		UpdateTask(task *models.Task) error
		DeleteTask(docID string, rev string) error
	}
)

// NewCouchDBRepository ...
func NewCouchDBRepository() CouchDBRepository {
	dataSourceName := viper.GetString("app.dataSourceName")
	client, err := kivik.New("couch", dataSourceName)
	if err != nil {
		fmt.Println("Failed to connect to couchdb: ", err.Error())
		return &couchdb{
			db: nil,
		}
	}

	dbname := viper.GetString("app.dbname")
	return &couchdb{
		db: client.DB(context.TODO(), dbname),
	}
}
