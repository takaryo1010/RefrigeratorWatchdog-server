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

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	imageRepository := repository.NewImageRepository()
	imageUsecase := usecase.NewImageUsecase(imageRepository)
	imageController := controller.NewImageController(imageUsecase)


	e := router.NewRouter(foodController, userController, imageController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
