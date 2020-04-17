package couchdb

import (
	"context"

	kivik "github.com/go-kivik/kivik/v3"
	"gitlab.com/renodesper/efishery-cli-app/models"
)

func (c *couchdb) GetTask(docID string) (*models.Task, error) {
	row := c.db.Get(context.TODO(), docID)

	var task models.Task
	if err := row.ScanDoc(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *couchdb) GetTasks(withDeleted bool) ([]*models.Task, error) {
	var docs *kivik.Rows
	var err error

	if !withDeleted {
		docs, err = c.db.Find(context.TODO(), map[string]interface{}{
			"selector": map[string]interface{}{
				"is_deleted": map[string]interface{}{
					"$ne": true,
				},
			},
		})
	} else {
		docs, err = c.db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})
	}

	if err != nil {
		return nil, err
	}

	var tasks []*models.Task

	for docs.Next() {
		var task models.Task
		if err := docs.ScanDoc(&task); err != nil {
			panic(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (c *couchdb) GetOutdatedTasks() ([]*models.Task, error) {
	localTasks, err := c.sqlite.GetTasks(true)
	if err != nil {
		return nil, err
	}

	remoteTasks, err := c.GetTasks(true)
	if err != nil {
		return nil, err
	}

	var newTasks []*models.Task

	for _, v1 := range localTasks {
		isDocExist := false

		for _, v2 := range remoteTasks {
			if v1.ID == v2.ID {
				if v1.CreatedAt.After(v2.CreatedAt) {
					newTasks = append(newTasks, v1)
				}

				isDocExist = true
				break
			}
		}

		if !isDocExist {
			newTasks = append(newTasks, v1)
		}
	}

	return newTasks, nil
}

func (c *couchdb) AddTask(task *models.Task) error {
	rev, err := c.db.Put(context.TODO(), task.ID, &task)
	if err != nil {
		return err
	}

	task.Rev = rev
	c.sqlite.AddTask(task)

	return nil
}

func (c *couchdb) UpdateTask(task *models.Task) error {
	newRev, err := c.db.Put(context.TODO(), task.ID, &task)
	if err != nil {
		return err
	}

	task.Rev = newRev
	c.sqlite.UpdateTask(task)

	return nil
}

func (c *couchdb) DeleteTask(task *models.Task) error {
	task.IsDeleted = true
	newRev, err := c.db.Put(context.TODO(), task.ID, &task)
	if err != nil {
		return err
	}

	task.Rev = newRev
	c.sqlite.DeleteTask(task)

	return nil
}

func (c *couchdb) DoneTask(task *models.Task) error {
	task.IsDone = true
	newRev, err := c.db.Put(context.TODO(), task.ID, &task)
	if err != nil {
		return err
	}

	task.Rev = newRev
	c.sqlite.DoneTask(task)

	return nil
}

func (c *couchdb) SyncTasks() error {
	newTasks, err := c.GetOutdatedTasks()
	if err != nil {
		return err
	}

	for _, v := range newTasks {
		err := c.UpdateTask(v)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
