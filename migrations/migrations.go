package migrations

import (
	"gorm.io/gorm"
	"simple-todo-list/pkg/entities"
)

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.Book{})
}
