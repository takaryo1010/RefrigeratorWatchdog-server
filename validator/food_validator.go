package validator

import (
	"RefrigeratorWatchdog-server/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IFoodValidator interface {
	ValidateFood(food model.Food) error
}

type foodValidator struct{}

func NewFoodValidator() IFoodValidator {
	return &foodValidator{}
}

func (fv *foodValidator) ValidateFood(food model.Food) error {
	if food.Tag == "" {
		food.Tag = "その他"
	}
	return validation.ValidateStruct(&food,
		validation.Field(&food.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&food.UserID, validation.Required),
		validation.Field(&food.OriginalCode, validation.Required, validation.Min(1), validation.Max(10000000000000)),
		validation.Field(&food.Quantity, validation.Required),
		validation.Field(&food.ExpirationDate, validation.Required),
		validation.Field(&food.ImageURL, validation.Required, validation.Length(1, 255)),
		validation.Field(&food.Memo, validation.Required, validation.Length(1, 1000)),
		validation.Field(&food.Tag, validation.Required, validation.In("野菜", "肉", "魚", "乳製品", "調味料", "卵", "飲料", "果物", "加工食品", "その他")),
	)
}
