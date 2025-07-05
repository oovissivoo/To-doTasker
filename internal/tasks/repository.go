package tasks

import "task-tracker/pkg/db"

type TaskRepository struct {
	Database *db.Db
}

func NewTaskRepository(database *db.Db) *TaskRepository {
	return &TaskRepository{
		Database: database,
	}
}

func (repo *TaskRepository) Create(task *Task) (*Task, error) {
	result := repo.Database.DB.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) Delete(id int) error {
	result := repo.Database.DB.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
