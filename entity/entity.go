package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	Email     string         `json:"email"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type Todo struct {
	ID              uint           `json:"id" gorm:"primarykey"`
	ActivityGroupID uint           `json:"activity_group_id"`
	Title           string         `json:"title"`
	IsActive        bool           `json:"is_active" gorm:"default:true"`
	Priority        string         `json:"priority" gorm:"default:'very-high'"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
