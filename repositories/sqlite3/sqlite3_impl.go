package sqlite3

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/renodesper/efishery-cli-app/models"
)

func (s *sqlite3) GetTask(docID string) (*models.Task, error) {
	dbname := viper.GetString("app.dbname")

	stmt, err := s.db.Prepare(fmt.Sprintf("SELECT _id, _rev, content, tags, created_at, is_deleted, is_done FROM %s WHERE _id = ?", dbname))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var task models.Task
	var createdAt string
	err = stmt.QueryRow(docID).Scan(&task.ID, &task.Rev, &task.Content, &task.Tags, &createdAt, &task.IsDeleted, &task.IsDone)
	if err != nil {
		panic(err)
	}

	t, err := time.Parse("2006-01-02T15:04:05.000000000-07:00", createdAt)
	if err != nil {
		return nil, err
	}

	task.CreatedAt = t
	return &task, nil
}

func (s *sqlite3) GetTasks() ([]*models.Task, error) {
	dbname := viper.GetString("app.dbname")
	rows, err := s.db.Query(fmt.Sprintf("SELECT _id, _rev, content, tags, created_at, is_deleted, is_done FROM %s WHERE is_deleted != 1", dbname))
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task

	defer rows.Close()
	for rows.Next() {
		var task models.Task
		var createdAt string

		err = rows.Scan(&task.ID, &task.Rev, &task.Content, &task.Tags, &createdAt, &task.IsDeleted, &task.IsDone)
		if err != nil {
			panic(err)
		}

		t, err := time.Parse("2006-01-02T15:04:05.000000000-07:00", createdAt)
		if err != nil {
			panic(err)
		}

		task.CreatedAt = t

		tasks = append(tasks, &task)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *sqlite3) AddTask(task *models.Task) error {
	dbname := viper.GetString("app.dbname")

	tx, _ := s.db.Begin()
	stmt, _ := tx.Prepare(fmt.Sprintf("INSERT INTO %s (_id, _rev, content, tags, created_at, is_deleted, is_done) VALUES (?,?,?,?,?,?,?)", dbname))
	_, err := stmt.Exec(task.ID, task.Rev, task.Content, task.Tags, task.CreatedAt, 0, 0)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (s *sqlite3) UpdateTask(task *models.Task) error {
	dbname := viper.GetString("app.dbname")
	isDeleted := 0
	isDone := 0

	if task.IsDeleted {
		isDeleted = 1
	}

	if task.IsDone {
		isDone = 1
	}

	tx, _ := s.db.Begin()
	q := fmt.Sprintf("UPDATE %s SET _rev=?,content=?,tags=?,created_at=?,is_deleted=?,is_done=? WHERE _id=?", dbname)
	stmt, _ := tx.Prepare(q)

	_, err := stmt.Exec(task.Rev, task.Content, task.Tags, task.CreatedAt, isDeleted, isDone, task.ID)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (s *sqlite3) DeleteTask(task *models.Task) error {
	dbname := viper.GetString("app.dbname")

	tx, _ := s.db.Begin()
	stmt, _ := tx.Prepare(fmt.Sprintf("UPDATE %s SET _rev=?,is_deleted=? WHERE _id=?", dbname))
	_, err := stmt.Exec(task.Rev, 1, task.ID)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (s *sqlite3) DoneTask(task *models.Task) error {
	dbname := viper.GetString("app.dbname")

	tx, _ := s.db.Begin()
	stmt, _ := tx.Prepare(fmt.Sprintf("UPDATE %s SET _rev=?,is_done=? WHERE _id=?", dbname))
	_, err := stmt.Exec(task.Rev, 1, task.ID)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
