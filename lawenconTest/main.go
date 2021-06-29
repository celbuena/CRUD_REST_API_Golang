package main

import (
	"lawenconTest/migration"
	"lawenconTest/utils"
	"lawenconTest/controller/router"
	"lawenconTest/controller/apihandler"
	"lawenconTest/usecase"
	"lawenconTest/adapter"
	"lawenconTest/db"
)

func main (){
	db := db.MYSQLConnection()
	migration.Migration()
	us := adapter.NewDataRepository(db)
	sr := usecase.NewServices(us)
	ct := apihandler.NewController(sr)
	rt := router.NewRoute(ct)

	utils.Start(rt.Mux())
}