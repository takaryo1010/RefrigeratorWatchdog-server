package controller

import (
	"RefrigeratorWatchdog-server/model"
	"RefrigeratorWatchdog-server/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IFoodController interface {
	GetFoodsByUserID(c echo.Context) error
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

// GetFoodsByUserID godoc
// @Summary Get foods by user id
// @Description Get foods by user id
// @ID get-foods-by-user-id
// @Accept  json
// @Produce  json
// @Param id path int true "UserID"
// @Success 200 {array} model.FoodResponse
// @Router /foods/{id} [get]
// @Tags foods
func (fc *foodController) GetFoodsByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	foods, err := fc.fu.GetFoodsByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, foods)
}


// CreateFood godoc
// @Summary Create food
// @Description Create food
// @ID create-food
// @Accept  json
// @Produce  json
// @Param food body model.FoodRequest true "Food"
// @Success 200 {object} model.FoodResponse
// @Router /foods [post]
// @Tags foods
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

// UpdateFood godoc
// @Summary Update food
// @Description Update food
// @ID update-food
// @Accept  json
// @Produce  json
// @Param id path int true "Food ID"
// @Param food body model.FoodRequest true "Food"
// @Success 200 {object} model.FoodResponse
// @Router /foods/{id} [put]
// @Tags foods
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

// DeleteFood godoc
// @Summary Delete food
// @Description Delete food
// @ID delete-food
// @Accept  json
// @Produce  json
// @Param id path int true "Food ID"
// @Success 200 {string} string "deleted"
// @Router /foods/{id} [delete]
// @Tags foods
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
