package migration

import (
	"fmt"

	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database"
	"github.com/anggriawanrilda88/myboilerplate/app/infrastructure/database/postgres/models"
)

// AutoMigratePostgres from model
func AutoMigratePostgres(DB *database.Database) (err error) {
	// role table migrate
	err = DB.AutoMigrate(&models.Role{})
	if err != nil {
		fmt.Println("failed to automigrate role model:", err.Error())
		return
	}

	// user table migrate
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("failed to automigrate user model:", err.Error())
		return
	}

	return err
}
