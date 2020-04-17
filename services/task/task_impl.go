package task

import (
	"fmt"
	"time"

	"gitlab.com/renodesper/efishery-cli-app/models"
)

func (t *task) GetTasks() ([]*models.Task, []*models.Task, error) {
	tasks, err := t.repository.GetTasks(false)
	if err != nil {
		return nil, nil, err
	}

	outdatedTasks, err := t.repository.GetOutdatedTasks()
	if err != nil {
		return nil, nil, err
	}

	return tasks, outdatedTasks, nil
}

func (t *task) AddTask(task *models.Task) error {
	task.CreatedAt = time.Now()
	err := t.repository.AddTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *task) UpdateTask(newTask *models.Task) error {
	task, err := t.repository.GetTask(newTask.ID)
	if err != nil {
		return err
	}

	if task == nil {
		fmt.Println("Cannot find the specified docID")
		return nil
	}

	if task.Content != newTask.Content {
		task.Content = newTask.Content
	}

	if task.Tags != newTask.Tags {
		task.Tags = newTask.Tags
	}

	task.CreatedAt = newTask.CreatedAt

	err = t.repository.UpdateTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *task) DeleteTask(docID string) error {
	task, err := t.repository.GetTask(docID)
	if err != nil {
		return err
	}

	err = t.repository.DeleteTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *task) DoneTask(docID string) error {
	task, err := t.repository.GetTask(docID)
	if err != nil {
		return err
	}

	err = t.repository.DoneTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *task) SyncTasks() error {
	err := t.repository.SyncTasks()
	if err != nil {
		return err
	}

	return nil
}
