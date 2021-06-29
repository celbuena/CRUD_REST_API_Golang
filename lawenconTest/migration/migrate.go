package migration

import (
	db2 "lawenconTest/db"
	"lawenconTest/domain"
)

func Migration() {
	db := db2.MYSQLConnection()
	db.AutoMigrate(&domain.Admin{})
}