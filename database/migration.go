package database

import (
	"fmt"
	"ptedi/models"
	"ptedi/pkg/connection"
)

func RunMigration() {
	err := connection.DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Failed! Create Table to Database")
	}
	fmt.Println("Create Table Success")
}
