package task

import (
	"fmt"

	"gitlab.com/renodesper/efishery-cli-app/models"
	"gitlab.com/renodesper/efishery-cli-app/repositories/couchdb"
	"gitlab.com/renodesper/efishery-cli-app/repositories/sqlite3"
)

type (
	// DBRepository ...
	DBRepository interface {
		GetTask(docID string) (*models.Task, error)
		GetTasks(withDeleted bool) ([]*models.Task, error)
		GetOutdatedTasks() ([]*models.Task, error)
		AddTask(task *models.Task) error
		UpdateTask(task *models.Task) error
		DeleteTask(task *models.Task) error
		DoneTask(task *models.Task) error
		SyncTasks() error
	}

	task struct {
		repository DBRepository
	}

	// TaskService ...
	TaskService interface {
		GetTasks() ([]*models.Task, []*models.Task, error)
		AddTask(task *models.Task) error
		UpdateTask(task *models.Task) error
		DeleteTask(docID string) error
		DoneTask(docID string) error
		SyncTasks() error
	}
)

// NewTaskService ...
func NewTaskService() TaskService {
	sqlite := sqlite3.NewSqlite3Repository()

	if couchdb.HasConnection() {
		return &task{
			repository: couchdb.NewCouchDBRepository(sqlite),
		}
	}

	fmt.Println("\n> No internet connection, using local data\n")
	return &task{
		repository: sqlite,
	}
}
