package couchdb

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v3"
	"github.com/spf13/viper"
	"gitlab.com/renodesper/efishery-cli-app/models"
	"gitlab.com/renodesper/efishery-cli-app/repositories/sqlite3"
)

type (
	couchdb struct {
		db     *kivik.DB
		sqlite sqlite3.Sqlite3Repository
	}

	// CouchDBRepository ...
	CouchDBRepository interface {
		GetTask(docID string) (*models.Task, error)
		GetTasks(withDeleted bool) ([]*models.Task, error)
		GetOutdatedTasks() ([]*models.Task, error)
		AddTask(task *models.Task) error
		UpdateTask(task *models.Task) error
		DeleteTask(task *models.Task) error
		DoneTask(task *models.Task) error
		SyncTasks() error
	}
)

// NewCouchDBRepository ...
func NewCouchDBRepository(sqlite sqlite3.Sqlite3Repository) CouchDBRepository {
	dataSourceName := viper.GetString("app.dataSourceName")
	client, err := kivik.New("couch", dataSourceName)
	if err != nil {
		fmt.Println("Failed to connect to couchdb: ", err.Error())
		panic(err)
	}

	dbname := viper.GetString("app.dbname")
	return &couchdb{
		db:     client.DB(context.TODO(), dbname),
		sqlite: sqlite,
	}
}

// HasConnection ...
func HasConnection() bool {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
