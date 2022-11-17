package todo

import (
	"errors"
	"testing-code/entity"

	"gorm.io/gorm"
)

type repositoryTodo struct {
	db *gorm.DB
}

type RepositoryTodo interface {
	Store(data entity.Todo) (*entity.Todo, error)
	Show(id uint) (*entity.Todo, error)
	Shows() ([]entity.Todo, error)
	Update(data *entity.Todo) (*entity.Todo, error)
	Delete(id uint) error
}

func NewRepositoryTodo(db *gorm.DB) *repositoryTodo {
	return &repositoryTodo{db}
}

func (rt *repositoryTodo) Store(data entity.Todo) (*entity.Todo, error) {
	err := rt.db.Debug().Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (rt *repositoryTodo) Shows() ([]entity.Todo, error) {
	var data []entity.Todo

	row := rt.db.Find(&data).RowsAffected
	if row <= 0 {
		return nil, errors.New("todo not found")
	}

	return data, nil
}

func (rt *repositoryTodo) Show(id uint) (*entity.Todo, error) {
	var data entity.Todo

	row := rt.db.Find(&data, id).RowsAffected
	if row != 1 {
		return nil, errors.New("todo not found")
	}

	return &data, nil
}

func (rt *repositoryTodo) Update(data *entity.Todo) (*entity.Todo, error) {
	row := rt.db.Debug().Updates(data).RowsAffected
	if row != 1 {
		return nil, errors.New("update todo failed")
	}

	return data, nil
}

func (rt *repositoryTodo) Delete(id uint) error {
	row := rt.db.Where("id = ?", id).Delete(&entity.Todo{}).RowsAffected
	if row != 1 {
		return errors.New("not found")
	}

	return nil
}
