package bootstrap

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-todo-list/config/yaml"
)

func DatabaseConnection(config *yaml.Config) (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(`root:@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local`, config.DB.Host, config.DB.Port, config.DB.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
