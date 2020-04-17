package task

import (
	"gitlab.com/renodesper/efishery-cli-app/models"
	"gitlab.com/renodesper/efishery-cli-app/repositories/couchdb"
)

type (
	task struct {
		repository couchdb.CouchDBRepository
	}

	// TaskService ...
	TaskService interface {
		GetTasks() ([]*models.Task, error)
		AddTask(task *models.Task) error
		UpdateTask(task *models.Task) error
		DeleteTask(rev string) error
	}
)

// NewTaskService ...
func NewTaskService() TaskService {
	return &task{
		repository: couchdb.NewCouchDBRepository(),
	}
}
