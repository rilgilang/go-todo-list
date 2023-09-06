package entities

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	ID        int        `gorm:"primary_key" json:"id"`
	UserID    int        `gorm:"type:int" json:"user_id"`
	Title     string     `gorm:"type:varchar(255)" json:"title"`
	Status    bool       `gorm:"default:false" json:"status"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

type DeleteTodoRequest struct {
	ID string `json:"id"`
}
