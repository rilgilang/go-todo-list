package migrations

import (
	"gorm.io/gorm"
	"simple-todo-list/internal/entities"
)

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&entities.Todo{}, &entities.User{})
}
