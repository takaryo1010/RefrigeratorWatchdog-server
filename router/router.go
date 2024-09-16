package router

import (
	"RefrigeratorWatchdog-server/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(fc controller.IFoodController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	f := e.Group("/foods")
	f.GET("/:id", fc.GetFoodByUserID)
	f.POST("", fc.CreateFood)
	f.PUT("/:id", fc.UpdateFood)
	f.DELETE("/:id", fc.DeleteFood)

	return e

}
