package entities

import (
	"gorm.io/gorm"
	"time"
)

// Book Constructs your Book model under entities.
type Book struct {
	gorm.Model
	ID        int        `gorm:"primary_key" json:"id"`
	Title     string     `gorm:"type:varchar(255)" json:"title"`
	Author    string     `gorm:"type:varchar(255)" json:"author"`
	CreatedAt time.Time  `gorm:"not null" json:"createdAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updatedAt" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

// DeleteRequest struct is used to parse Delete Requests for Books
type DeleteRequest struct {
	ID string `json:"id"`
}
