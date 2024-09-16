package main

import (
	"RefrigeratorWatchdog-server/controller"
	"RefrigeratorWatchdog-server/db"
	"RefrigeratorWatchdog-server/repository"
	"RefrigeratorWatchdog-server/router"
	"RefrigeratorWatchdog-server/usecase"
	"RefrigeratorWatchdog-server/validator"
	"fmt"
	"os"
)

func main() {
	db := db.NewDB()
	foodValidator := validator.NewFoodValidator()
	foodRepository := repository.NewFoodRepository(db)
	foodUsecase := usecase.NewFoodUsecase(foodRepository, foodValidator)
	foodController := controller.NewFoodController(foodUsecase)
	e := router.NewRouter(foodController)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
