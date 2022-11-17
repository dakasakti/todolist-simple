package activity

import (
	"errors"
	"testing-code/entity"

	"gorm.io/gorm"
)

type repositoryActivity struct {
	db *gorm.DB
}

type RepositoryActivity interface {
	Store(data entity.Activity) (*entity.Activity, error)
	Show(id uint) (*entity.Activity, error)
	Shows() ([]entity.Activity, error)
	Update(data *entity.Activity) (*entity.Activity, error)
	Delete(id uint) error
}

func NewRepositoryActivity(db *gorm.DB) *repositoryActivity {
	return &repositoryActivity{db}
}

func (ra *repositoryActivity) Store(data entity.Activity) (*entity.Activity, error) {
	err := ra.db.Create(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (ra *repositoryActivity) Shows() ([]entity.Activity, error) {
	var data []entity.Activity

	row := ra.db.Find(&data).RowsAffected
	if row <= 0 {
		return nil, errors.New("activity not found")
	}

	return data, nil
}

func (ra *repositoryActivity) Show(id uint) (*entity.Activity, error) {
	var data entity.Activity

	row := ra.db.Find(&data, id).RowsAffected
	if row != 1 {
		return nil, errors.New("activity not found")
	}

	return &data, nil
}

func (ra *repositoryActivity) Update(data *entity.Activity) (*entity.Activity, error) {
	row := ra.db.Debug().Updates(data).RowsAffected
	if row != 1 {
		return nil, errors.New("update activity failed")
	}

	return data, nil
}

func (ra *repositoryActivity) Delete(id uint) error {
	row := ra.db.Where("id = ?", id).Delete(&entity.Activity{}).RowsAffected
	if row != 1 {
		return errors.New("not found")
	}

	return nil
}
