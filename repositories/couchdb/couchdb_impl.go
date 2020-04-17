package couchdb

import (
	"context"

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

func (c *couchdb) GetTasks() ([]*models.Task, error) {
	docs, err := c.db.Find(context.TODO(), map[string]interface{}{
		"selector": map[string]interface{}{
			"is_deleted": map[string]interface{}{
				"$ne": true,
			},
		},
	})
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
