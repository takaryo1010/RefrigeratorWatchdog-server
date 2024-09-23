package router

import (
	"RefrigeratorWatchdog-server/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @
func NewRouter(fc controller.IFoodController, uc controller.IUserController ,ic controller.IImageController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	f := e.Group("/foods")
	f.GET("/:id", fc.GetFoodsByUserID)
	f.POST("", fc.CreateFood)
	f.PUT("/:id", fc.UpdateFood)
	f.DELETE("/:id", fc.DeleteFood)

	//POST例
	/*
	   	{
	     "name": "オレンジ",
	     "user_id": 1,
	     "original_code": 12456456,
	     "quantity": 5,
	     "expiration_date": "2024-12-15T00:00:00Z",
	     "image_url": "https://example.com/images/orange.jpg",
	     "memo": "新鮮なオレンジだったものです"
	   }
	*/

	u := e.Group("/users")
	u.GET("/:email", uc.GetUser)
	u.POST("", uc.CreateUser)
	u.PUT("/:email", uc.UpdateUser)
	// POSTする際にuser情報をすべて送信する必要がある
	u.DELETE("", uc.DeleteUser)

	//POST例
	/*

	   {
	     "username": "山田太郎",
	     "email": "sample@gmail.com",
	     "password": "password"
	   }
	*/
	i := e.Group("/images")
	i.GET("/:imageURL", ic.FetchImage)
	i.POST("", ic.UploadImage)

	return e

}
