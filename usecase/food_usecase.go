package usecase

import (
	"RefrigeratorWatchdog-server/model"

	"RefrigeratorWatchdog-server/repository"
	"RefrigeratorWatchdog-server/validator"
)

type IFoodUsecase interface {
	GetFoodByUserID(userID uint) (model.FoodResponse, error)
	CreateFood(food model.Food) (model.FoodResponse, error)
	UpdateFood(food model.Food, id uint) (model.FoodResponse, error)
	DeleteFood(id uint) error
}

type foodUsecase struct {
	fr repository.IFoodRepository
	fv validator.IFoodValidator
}

func NewFoodUsecase(fr repository.IFoodRepository, fv validator.IFoodValidator) IFoodUsecase {
	return &foodUsecase{fr, fv}
}

func (fu *foodUsecase) GetFoodByUserID(userID uint) (model.FoodResponse, error) {
	food := model.Food{}
	if err := fu.fr.GetFoodByUserID(&food, userID); err != nil {
		return model.FoodResponse{}, err
	}
	return model.FoodResponse{
		ID:             food.ID,
		Name:           food.Name,
		UserID:         food.UserID,
		OriginalCode:   food.OriginalCode,
		Quantity:       food.Quantity,
		CreatedAt:      food.CreatedAt,
		ExpirationDate: food.ExpirationDate,
		ImageURL:       food.ImageURL,
		Memo:           food.Memo,
	}, nil
}

func (fu *foodUsecase) CreateFood(food model.Food) (model.FoodResponse, error) {
	if err := fu.fv.ValidateFood(food); err != nil {
		return model.FoodResponse{}, err
	}

	if err := fu.fr.CreateFood(&food); err != nil {
		return model.FoodResponse{}, err
	}

	return model.FoodResponse{
		ID:             food.ID,
		Name:           food.Name,
		UserID:         food.UserID,
		OriginalCode:   food.OriginalCode,
		Quantity:       food.Quantity,
		CreatedAt:      food.CreatedAt,
		ExpirationDate: food.ExpirationDate,
		ImageURL:       food.ImageURL,
		Memo:           food.Memo,
	}, nil
}

func (fu *foodUsecase) UpdateFood(food model.Food, id uint) (model.FoodResponse, error) {
	if err := fu.fv.ValidateFood(food); err != nil {
		return model.FoodResponse{}, err
	}

	if err := fu.fr.UpdateFood(&food, id); err != nil {
		return model.FoodResponse{}, err
	}

	return model.FoodResponse{
		ID:             food.ID,
		Name:           food.Name,
		UserID:         food.UserID,
		OriginalCode:   food.OriginalCode,
		Quantity:       food.Quantity,
		CreatedAt:      food.CreatedAt,
		ExpirationDate: food.ExpirationDate,
		ImageURL:       food.ImageURL,
		Memo:           food.Memo,
	}, nil
}

func (fu *foodUsecase) DeleteFood(id uint) error {

	if err := fu.fr.DeleteFood(id); err != nil {
		return err
	}

	return nil
}
