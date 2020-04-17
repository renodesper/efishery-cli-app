package couchdb

import (
	"context"

	kivik "github.com/go-kivik/kivik/v3"
	"gitlab.com/renodesper/efishery-cli-app/models"
)

func (c *couchdb) GetTasks() (*kivik.Rows, error) {
	docs, err := c.db.AllDocs(context.TODO(), kivik.Options{"include_docs": true})
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (c *couchdb) AddTask(task *models.Task) error {
	rev, err := c.db.Put(context.TODO(), task.ID, &task)
	if err != nil {
		return err
	}

	task.Rev = rev
	return nil
}

func (c *couchdb) UpdateTask(task *models.Task) error {
	newRev, err := c.db.Put(context.TODO(), task.ID, &task)
	if err != nil {
		return err
	}

	task.Rev = newRev
	return nil
}

func (c *couchdb) DeleteTask(docID string, rev string) error {
	_, err := c.db.Delete(context.TODO(), docID, rev)
	if err != nil {
		return err
	}

	return nil
}
