package main

import (
	"RefrigeratorWatchdog-server/db"
	"RefrigeratorWatchdog-server/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
	dbConn.AutoMigrate(&model.Food{})
}
