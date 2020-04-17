package task

import (
	"gitlab.com/renodesper/efishery-cli-app/models"
)

func (t *task) GetTasks() ([]*models.Task, error) {
	docs, err := t.repository.GetTasks()
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

func (t *task) AddTask(task *models.Task) error {
	err := t.repository.AddTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *task) UpdateTask(task *models.Task) error {
	err := t.repository.UpdateTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *task) DeleteTask(rev string) error {
	return nil
}
