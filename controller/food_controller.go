package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type IFoodController interface {
	GetFoodByUserID(c echo.Context) error
	CreateFood(c echo.Context) error
	UpdateFood(c echo.Context) error
	DeleteFood(c echo.Context) error
}
type foodController struct {
	fu usecase.IFoodUsecase
}

func NewFoodController(fu usecase.IFoodUsecase) IFoodController {
	return &foodController{fu}
}

func (fc *foodController) GetFoodByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	food, err := fc.fu.GetFoodByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, food)
}

func (fc *foodController) CreateFood(c echo.Context) error {
	food := model.Food{}
	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	createdFood, err := fc.fu.CreateFood(food)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, createdFood)
}

func (fc *foodController) UpdateFood(c echo.Context) error {
	food := model.Food{}
	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	updatedFood, err := fc.fu.UpdateFood(food, uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, updatedFood)
}

func (fc *foodController) DeleteFood(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = fc.fu.DeleteFood(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "deleted")
}

