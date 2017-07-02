package models

import "github.com/bturbes/demo_api/database"

// MigrateAll runs all the migrations for the models
func MigrateAll() {
	database.Db.AutoMigrate(&Article{})
}
