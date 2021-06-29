package db

import (
	"fmt"
	"lawenconTest/utils"
	"gorm.io/gorm"
	_ "github.com/go-sql-driver/mysql"

)

func MYSQLConnection() *gorm.DB {
	dataConfig, err := utils.GetConfigFile()
	if err != nil {
		return nil
	}

	connInfo := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, dataConfig.Database.Username, dataConfig.Database.Password, dataConfig.Database.Host, dataConfig.Database.Port, dataConfig.Database.Name)

	db, err := gorm.Open("mysql", connInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
