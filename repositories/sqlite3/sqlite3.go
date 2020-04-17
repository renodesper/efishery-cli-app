package sqlite3

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"gitlab.com/renodesper/efishery-cli-app/models"
)

type (
	sqlite3 struct {
		db *sql.DB
	}

	// Sqlite3Repository ...
	Sqlite3Repository interface {
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

// NewSqlite3Repository ...
func NewSqlite3Repository() Sqlite3Repository {
	dbname := viper.GetString("app.dbname")

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s.db", dbname))
	if err != nil {
		panic(err)
	}

	db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (_id TEXT, _rev TEXT, content TEXT, tags TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, is_deleted INT DEFAULT 0, is_done INT DEFAULT 0)", dbname))

	return &sqlite3{
		db: db,
	}
}
