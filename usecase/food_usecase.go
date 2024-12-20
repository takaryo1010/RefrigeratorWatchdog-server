package usecase

import (
	"RefrigeratorWatchdog-server/model"

	"RefrigeratorWatchdog-server/repository"
	"RefrigeratorWatchdog-server/validator"
)

type IFoodUsecase interface {
	GetFoodsByUserID(userID uint) ([]model.FoodResponse, error)
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

func (fu *foodUsecase) GetFoodsByUserID(userID uint) ([]model.FoodResponse, error) {
	foods := []model.Food{}
	if err := fu.fr.GetFoodsByUserID(&foods, userID); err != nil {
		return nil, err
	}
	resFoods := []model.FoodResponse{}
	for _, food := range foods {
		resFoods = append(resFoods, model.FoodResponse{
			ID:             food.ID,
			Name:           food.Name,
			UserID:         food.UserID,
			OriginalCode:   food.OriginalCode,
			Quantity:       food.Quantity,
			CreatedAt:      food.CreatedAt,
			ExpirationDate: food.ExpirationDate,
			ImageURL:       food.ImageURL,
			Tag:            food.Tag,
			Memo:           food.Memo,
		})
	}
	return resFoods, nil
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
		Tag:            food.Tag,
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
		Tag:            food.Tag,
		Memo:           food.Memo,
	}, nil
}

func (fu *foodUsecase) DeleteFood(id uint) error {

	if err := fu.fr.DeleteFood(id); err != nil {
		return err
	}

	return nil
}
